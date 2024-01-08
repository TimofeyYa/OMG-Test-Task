package repository

import (
	"context"
	"omg/intermal/repository/api"

	"github.com/jackc/pgx/v4/pgxpool"
)

type YClients interface {
	GetAPIStaff(c context.Context, companyId int) (*any, error)
}

type Store interface {
	CreateTask(c context.Context) (int, error)
	PerformActiveTasks(c context.Context) error
	GetStoreStaff(c context.Context, taskId int) (*any, error)
}

type Repository struct {
	YClients
	Store
}

func NewRepository(apiUri string, pool *pgxpool.Pool) *Repository {
	return &Repository{
		YClients: api.NewYClients(apiUri),
	}
}
