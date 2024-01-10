package service

import (
	"context"
	"omg/intermal/models"
	"omg/intermal/repository"
)

type Service interface {
	CreateTask(c context.Context, companyId int) (int, error)
	GetStatusTask(c context.Context, taskId int) (*models.Status, error)
	GetStaff(c context.Context, taskId int) (*any, error)
	ResolveTasks(c context.Context) error
}

type service struct {
	repo *repository.Repository
}

func NewService(r *repository.Repository) Service {
	return &service{
		repo: r,
	}
}
