package services

import (
	"context"
	"otus_highload/internal/models"
	"otus_highload/internal/services/auth"
	"otus_highload/internal/services/users"
	"otus_highload/internal/storages"
)

func NewService(storage *storages.Storage) *Service {
	return &Service{
		Auth:  auth.NewAuthService(storage),
		Users: users.NewUsersService(storage),
	}
}

type Service struct {
	Auth
	Users
}

type Auth interface {
	SignIn(ctx context.Context, dto auth.SignInDto) (string, bool, error)
	SignUp(ctx context.Context, dto auth.SignUpDto) error
}

type Users interface {
	UserById(ctx context.Context, id int64) (*models.User, error)
}
