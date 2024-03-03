package signup

import (
	"context"
	"errors"
	"otus_highload/internal/domain"
	"otus_highload/internal/lib/errs"
	"time"
)

type (
	UserReader interface {
		UserByEmail(ctx context.Context, email string) (*domain.User, error)
	}
	UsersReadr interface {
		Add(ctx context.Context, u *domain.User) error
	}
	jwt interface {
		NewUserToken(user domain.User) (string, error)
	}
)

type Service struct {
	userReader UserReader
	userWriter UsersReadr
	jwt        jwt
}

func New(userReader UserReader, userWriter UsersReadr, jwt jwt) *Service {
	return &Service{userReader: userReader, userWriter: userWriter, jwt: jwt}
}

func (s Service) SignUp(ctx context.Context, user *domain.User) (err error) {
	_, err = s.userReader.UserByEmail(ctx, user.Email)
	if err == nil {
		return errs.ErrUserWithEmailAlreadyExists{Email: user.Email}
	}
	if errors.As(err, &errs.ErrUserByEmailNotFound{}) {
		if err = user.GeneratePasswordHash(); err != nil {
			return err
		}
		user.CreatedAt = time.Now()
		return s.userWriter.Add(ctx, user)
	}
	return err
}
