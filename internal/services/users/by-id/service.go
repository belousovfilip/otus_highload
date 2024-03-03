package byid

import (
	"context"
	"errors"
	"fmt"
	"otus_highload/internal/domain"
	"otus_highload/internal/lib/errs"
)

type (
	UsersReader interface {
		UserByID(context.Context, int64) (*domain.User, error)
	}
)

type Service struct {
	userReader UsersReader
}

func New(userReader UsersReader) *Service {
	return &Service{userReader: userReader}
}

func (s Service) UserByID(ctx context.Context, id int64) (*domain.User, error) {
	user, err := s.userReader.UserByID(ctx, id)
	if err != nil {
		if errors.As(err, &errs.ErrUserByIDNotFound{}) {
			return nil, err
		}
		return nil, fmt.Errorf("service:users:user-by-id: %w", err)
	}
	return user, nil
}
