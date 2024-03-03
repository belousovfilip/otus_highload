package errs

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrUserPasswordIsNotEqual = errors.New("user password is not equal")

	ErrInvalidEmail     = errors.New("invalid email")
	ErrInvalidPassword  = errors.New("invalid password")
	ErrInvalidFirstName = errors.New("invalid first name")
	ErrInvalidLastName  = errors.New("invalid last name")
)

type (
	ErrUserByIDNotFound struct {
		ID int64
	}
	ErrUserByEmailNotFound struct {
		Email string
	}
	ErrUserWithEmailAlreadyExists struct {
		Email string
	}
)

func (r ErrUserByIDNotFound) Error() string {
	return fmt.Sprintf("user by id %d not found", r.ID)
}

func (r ErrUserByEmailNotFound) Error() string {
	return fmt.Sprintf("user by email %s not found", r.Email)
}

func (r ErrUserWithEmailAlreadyExists) Error() string {
	return fmt.Sprintf("user with email %s already exists", r.Email)
}

type ValidationErrors map[string]string

func (v ValidationErrors) Error() string {
	b := strings.Builder{}
	for field, value := range v {
		b.WriteString(field)
		b.WriteString(": ")
		b.WriteString(value)
		b.WriteString(";")
	}
	return b.String()
}
