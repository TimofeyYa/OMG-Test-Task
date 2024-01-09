package store

import (
	"context"
	"fmt"
	"omg/intermal/models"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const InProgressStatus = "InProgress"

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
	if err := s.conn.QueryRow(c, sql, companyId, InProgressStatus).Scan(&taskId); err != nil {
		return 0, err
	}

	return taskId, nil
}

func (s *Store) GetTaskStatus(c context.Context, taskId int) (string, error) {
	return "", nil
}

func (s *Store) PerformActiveTasks(c context.Context) error {
	tx, err := s.conn.BeginTx(c, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		return fmt.Errorf("err create tx for perform tasks: %s", err)
	}

	sqlSelecActive := `SELECT id, company_id FROM tasks WHERE status = $1`

	rows, err := tx.Query(c, sqlSelecActive, InProgressStatus)
	if err != nil {
		return fmt.Errorf("err get active tasks: %s", err)
	}

	for rows.Next() {
		var task models.Task
		rows.Scan(&task.Id, &task.CompanyId)
		fmt.Println(task)
	}

	return nil
}

func (s *Store) GetStoreStaff(c context.Context, taskId int) (*any, error) {
	return nil, nil
}
