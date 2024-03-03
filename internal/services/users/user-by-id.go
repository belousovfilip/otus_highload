package users

import (
	"context"
	"otus_highload/internal/models"
)

func (s *Service) UserById(ctx context.Context, id int64) (*models.User, error) {
	user, err := s.storage.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
