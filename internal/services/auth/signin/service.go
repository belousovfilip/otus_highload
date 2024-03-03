package signin

import (
	"context"
	"errors"
	"otus_highload/internal/domain"
	"otus_highload/internal/lib/errs"
)

type (
	UserReader interface {
		UserByEmail(ctx context.Context, email string) (*domain.User, error)
	}
	Jwt interface {
		NewUserToken(user domain.User) (string, error)
	}
)

type Service struct {
	usersReader UserReader
	jwt         Jwt
}

func New(usersReader UserReader, jwt Jwt) *Service {
	return &Service{usersReader: usersReader, jwt: jwt}
}

func (s Service) SignIn(ctx context.Context, email, password string) (token string, err error) {
	var user *domain.User
	user, err = s.usersReader.UserByEmail(ctx, email)
	if err != nil {
		if errors.As(err, &errs.ErrUserByEmailNotFound{}) {
			return "", err
		}
		return "", err
	}
	if err = user.ComparePassword(password); err != nil {
		return "", err
	}
	if token, err = s.jwt.NewUserToken(*user); err != nil {
		return "", err
	}
	return token, nil
}
