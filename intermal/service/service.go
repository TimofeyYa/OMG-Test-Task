package service

import "omg/intermal/repository"

type Service interface {
}

type service struct {
	repo *repository.Repository
}

func NewService(r *repository.Repository) Service {
	return &service{
		repo: r,
	}
}
