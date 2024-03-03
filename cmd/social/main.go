package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"net/http"
	"os"
	"otus_highload/internal/handlers"
	c "otus_highload/internal/lib/config"
	l "otus_highload/internal/lib/logger"
	"otus_highload/internal/services"
	"otus_highload/internal/storages"
)

func main() {
	logger := l.Init()

	logger.Info("config initializing")
	config, err := NewConfig()
	handleError(logger, err)

	logger.Info("database initializing")
	db, err := NewDB(config.Database)
	handleError(logger, err)
	defer db.Close()

	logger.Info("storages initializing")
	storage := storages.NewStorage(db)
	logger.Info("services initializing")
	service := services.NewService(storage)
	logger.Info("handlers initializing")
	handler := handlers.NewHandler(service, logger).InitHandlers()

	server := http.Server{
		Addr:    config.HTTPServer.Host + ":" + config.HTTPServer.Port,
		Handler: handler,
	}
	logger.Info("server initializing http://" + server.Addr)
	handleError(logger, server.ListenAndServe())
}

func handleError(logger *slog.Logger, err error) {
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func NewDB(c c.Database) (*sqlx.DB, error) {
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.User, c.Password, c.Host, c.Port, c.Name)
	return sqlx.Open("mysql", dbSource)
}

func NewConfig() (*c.Config, error) {
	cPath := os.Getenv("CONFIG_PATH")
	if cPath == "" {
		cPath = "config/config.yaml"
	}
	return c.Init(cPath)
}
