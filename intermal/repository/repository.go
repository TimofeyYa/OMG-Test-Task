package repository

import (
	"context"
	"omg/intermal/models"
	"omg/intermal/repository/store"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store interface {
	CreateTask(c context.Context, companyId int) (int, error)
	GetTaskStatus(c context.Context, taskId int) (*models.Status, error)
	PerformActiveTasks(c context.Context) error
	GetStoreStaff(c context.Context, taskId int) (*any, error)
}

type Repository struct {
	Store
}

func NewRepository(apiUri string, pool *pgxpool.Pool, tokens *models.TokenPair) *Repository {
	return &Repository{
		Store: store.NewStore(pool, apiUri, tokens),
	}
}
