package auth

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"otus_highload/internal/lib/appErrors"
	"otus_highload/internal/lib/jwt"
	"time"
)

type SignInDto struct {
	Email    string `validate:"required,min=3,max=255"`
	Password string `validate:"required,min=3,max=255"`
}

func (s *Service) SignIn(ctx context.Context, dto SignInDto) (token string, isAuth bool, err error) {
	if err := validator.New().Struct(dto); err != nil {
		return "", false, appErrors.NewValidationErrorsFromValidator(err.(validator.ValidationErrors))
	}
	user, err := s.storage.FindUserByEmail(ctx, dto.Email)
	if err != nil {
		return "", false, err
	}
	if user == nil {
		return "", false, nil
	}
	isEqual, err := user.ComparePassword(dto.Password)
	if err != nil {
		return "", false, fmt.Errorf("compare password; %s", err)
	}
	if !isEqual {
		return "", false, nil
	}
	token, err = jwt.NewToken(*user, time.Hour*24)
	if err != nil {
		return "", false, err
	}
	return token, true, nil
}
