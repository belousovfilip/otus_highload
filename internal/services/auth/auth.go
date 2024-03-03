package auth

import (
	"otus_highload/internal/storages"
)

func NewAuthService(s *storages.Storage) *Service {
	return &Service{storage: s}
}

type Service struct {
	storage *storages.Storage
}
