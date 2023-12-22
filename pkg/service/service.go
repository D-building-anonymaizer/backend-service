package service

import "github.com/D-building-anonymaizer/backend-service/pkg/repository"

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
