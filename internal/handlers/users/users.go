package users

import (
	"log/slog"
	"otus_highload/internal/models"
	"otus_highload/internal/services"
	"time"
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

type UserResource struct {
	Id        int64     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Age       int       `json:"age,omitempty"`
	Gender    string    `json:"gender"`
	City      string    `json:"city"`
	Interests string    `json:"interests"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewUserResource(u *models.User) UserResource {
	return UserResource{
		Id:        u.Id,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Age:       u.Age,
		Gender:    u.Gender,
		City:      u.City,
		Interests: u.Interests,
		CreatedAt: u.CreatedAt,
	}
}
