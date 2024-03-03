package auth

import (
	"log/slog"
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
