package auth

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"otus_highload/internal/lib/appErrors"
	"otus_highload/internal/models"
	"time"
)

type SignUpDto struct {
	FirstName string `validate:"required,min=3,max=255"`
	LastName  string `validate:"required,min=3,max=255"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required,min=5,max=100"`
}

func (s *Service) SignUp(ctx context.Context, dto SignUpDto) error {
	if err := validator.New().Struct(dto); err != nil {
		return appErrors.NewValidationErrorsFromValidator(err.(validator.ValidationErrors))
	}
	user, err := s.storage.FindUserByEmail(ctx, dto.Email)
	fmt.Println(user)
	if err != nil {
		return err
	}
	if user != nil {
		errs := appErrors.NewValidationErrors()
		errs["email"] = "email exists"
		return errs
	}
	user = &models.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
	}
	if err := user.GeneratePasswordHash(dto.Password); err != nil {
		return err
	}
	user.CreatedAt = time.Now()
	return s.storage.AddUser(ctx, user)
}
