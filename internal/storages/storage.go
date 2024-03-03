package storages

import (
	"context"
	"github.com/jmoiron/sqlx"
	"otus_highload/internal/models"
	"otus_highload/internal/storages/mysql/users"
)

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		users.NewUserStorage(db),
	}
}

type Storage struct {
	User
}

type User interface {
	AddUser(ctx context.Context, u *models.User) error
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	FindUserById(ctx context.Context, id int64) (*models.User, error)
}
