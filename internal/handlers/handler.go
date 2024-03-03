package handlers

import (
	"log/slog"
	"net/http"
	"otus_highload/internal/handlers/auth"
	"otus_highload/internal/handlers/users"
	"otus_highload/internal/services"
)

func NewHandler(s *services.Service, log *slog.Logger) *Handler {
	return &Handler{
		service: s,
		log:     log,
	}
}

type Handler struct {
	service *services.Service
	log     *slog.Logger
}

func (h *Handler) InitHandlers() *http.ServeMux {
	s := http.NewServeMux()

	authH := auth.NewHandler(h.service, h.log)
	s.HandleFunc("POST /api/auth/signup", authH.SignUp)
	s.HandleFunc("POST /api/auth/signin", authH.SignIn)

	usersH := users.NewHandler(h.service, h.log)
	s.HandleFunc("GET /api/users/{id}", usersH.UserById)
	return s
}
