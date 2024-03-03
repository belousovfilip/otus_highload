package domain

import (
	"errors"
	"otus_highload/internal/lib/errs"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
	Password  string
	Age       int
	Gender    string
	City      string
	Interests string
	CreatedAt time.Time `db:"created_at"`
}

func (u *User) GeneratePasswordHash() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) ComparePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err == nil {
		return nil
	}
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return errs.ErrUserPasswordIsNotEqual
	}
	return err
}
