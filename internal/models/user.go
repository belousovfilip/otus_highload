package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id        int64
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

func (u *User) GeneratePasswordHash(p string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) ComparePassword(p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err == nil {
		return true, nil
	}
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false, nil
	}
	return false, err
}
