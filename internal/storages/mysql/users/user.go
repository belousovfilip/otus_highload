package users

import (
	"github.com/jmoiron/sqlx"
)

func NewUserStorage(db *sqlx.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

type UserStorage struct {
	db *sqlx.DB
}
