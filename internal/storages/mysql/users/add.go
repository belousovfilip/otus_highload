package users

import (
	"context"
	"otus_highload/internal/models"
)

func (s UserStorage) AddUser(ctx context.Context, u *models.User) error {
	q := "INSERT INTO users (email, password, created_at) VALUE (:email, :password, :created_at)"
	if _, err := s.db.NamedExecContext(ctx, q, u); err != nil {
		return err
	}
	return nil
}
