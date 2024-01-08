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

func (s *Store) CreateTask(c context.Context, companyId int) (int, error) {
	sql := `INSERT INTO tasks (company_id, status) VALUES ($1, $2) RETURNING "id"`

	var taskId int
	if err := s.conn.QueryRow(c, sql, companyId, "InProgress").Scan(&taskId); err != nil {
		return 0, err
	}

	return taskId, nil
}

func (s *Store) PerformActiveTasks(c context.Context) error {
	return nil
}

func (s *Store) GetStoreStaff(c context.Context, taskId int) (*any, error) {
	return nil, nil
}
