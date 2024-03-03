package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	signin_handler "otus_highload/internal/handlers/auth/signin"
	signup_handler "otus_highload/internal/handlers/auth/signup"
	user_by_id_service "otus_highload/internal/handlers/users/byid"
	config_lib "otus_highload/internal/lib/config"
	db_lib "otus_highload/internal/lib/db"
	jwt_lib "otus_highload/internal/lib/jwt"
	logger_lib "otus_highload/internal/lib/logger"
	http_response_lib "otus_highload/internal/lib/response"
	signin_service "otus_highload/internal/services/auth/signin"
	signup_service "otus_highload/internal/services/auth/signup"
	users_storage "otus_highload/internal/storages/mysql/users"
)

func main() {
	logger := logger_lib.Init()
	ctx := context.Background()
	cfg, err := config_lib.InitAppConfigFromEnv()
	handleError(logger, err)

	pool := db_lib.NewPool()
	err = pool.Connect(ctx, cfg.DBPrimary.Connection)
	handleError(logger, err)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	usersReader := users_storage.NewReader(pool)
	usersWriter := users_storage.NewWriter(pool)
	resp := http_response_lib.NewJSON(logger)
	jwt := jwt_lib.New()

	router.Post(
		"/api/auth/signin",
		signin_handler.New(logger, resp, signin_service.New(usersReader, jwt)).ServeHTTP,
	)
	router.Post(
		"/api/auth/signup",
		signup_handler.New(logger, resp, signup_service.New(usersReader, usersWriter, jwt)).ServeHTTP,
	)
	router.Get(
		"/api/users/{id}",
		user_by_id_service.NewHandler(logger, resp, usersReader).ServeHTTP,
	)
	server := &http.Server{
		Addr:              cfg.HTTPServer.Addr,
		Handler:           router,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
		ReadTimeout:       3 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		return
	}
}

func handleError(logger *slog.Logger, err error) {
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
