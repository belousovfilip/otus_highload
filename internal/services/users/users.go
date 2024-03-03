package users

import (
	"otus_highload/internal/storages"
)

func NewUsersService(s *storages.Storage) *Service {
	return &Service{storage: s}
}

type Service struct {
	storage *storages.Storage
}
