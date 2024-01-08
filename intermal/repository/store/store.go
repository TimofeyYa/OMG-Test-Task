package store

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	conn   *pgxpool.Pool
	apiUri string
}

func NewStore(conn *pgxpool.Pool, apiUri string) *Store {
	return &Store{
		conn:   conn,
		apiUri: apiUri,
	}
}

func (s *Store) CreateTask(c context.Context) (int, error) {
	return 0, nil
}

func (s *Store) PerformActiveTasks(c context.Context) error {
	return nil
}

func (s *Store) GetStoreStaff(c context.Context, taskId int) (*any, error) {
	return nil, nil
}
