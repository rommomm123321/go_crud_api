package service

import "github.com/rommomm123321/go_crud_api/pkg/repository"

type Service struct {
}

func NewService(r *repository.Repository) *Service {
	return &Service{}
}
