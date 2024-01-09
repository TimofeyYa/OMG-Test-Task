package store

import (
	"context"
	"fmt"
	"omg/intermal/models"
	"sync"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const InProgressStatus = "InProgress"
const DoneStatus = "Done"
const ErrStatus = "Error"

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

type StaffAPIItem struct {
	TaskId    int
	StaffData any
	Err       error
}

func (s *Store) PerformActiveTasks(c context.Context) error {
	tx, err := s.conn.BeginTx(c, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		return fmt.Errorf("err create tx for perform tasks: %s", err)
	}

	sqlSelecActive := `SELECT id, company_id FROM tasks WHERE status = $1 FOR UPDATE`

	rows, err := tx.Query(c, sqlSelecActive, InProgressStatus)
	if err != nil {
		return fmt.Errorf("err get active tasks: %s", err)
	}

	wg := &sync.WaitGroup{}

	staffDataCh := make(chan StaffAPIItem)
	for rows.Next() {
		var task models.Task
		err = rows.Scan(&task.Id, &task.CompanyId)
		if err != nil {
			return fmt.Errorf("err scan tasks: %s", err)
		}

		wg.Add(1)
		go func(task models.Task) {
			data, err := s.GetStaffFromApi(c, task.CompanyId)
			staffDataCh <- StaffAPIItem{
				StaffData: data,
				Err:       err,
				TaskId:    task.Id,
			}
			wg.Done()
		}(task)

		fmt.Println(task)
	}

	go func() {
		wg.Wait()
		close(staffDataCh)
	}()

	sqlAddResult := `INSERT INTO tasks_result (task_id, result) VALUES ($1, $2)`
	sqlUpdateStatus := `UPDATE tasks t SET status  = $1 WHERE id = $2`
	for staff := range staffDataCh {
		if staff.Err != nil {
			_, err = tx.Exec(c, sqlUpdateStatus, ErrStatus, staff.TaskId)
			if err != nil {
				return fmt.Errorf("err update status: %s", err)
			}
			continue
		}
		_, err := tx.Exec(c, sqlAddResult, staff.TaskId, staff.StaffData)
		if err != nil {
			return fmt.Errorf("err add result: %s", err)
		}

		_, err = tx.Exec(c, sqlUpdateStatus, DoneStatus, staff.TaskId)
		if err != nil {
			return fmt.Errorf("err update status: %s", err)
		}
		fmt.Println(staff.StaffData)
	}

	fmt.Println("ALL")
	tx.Commit(c)

	return nil
}

func (s *Store) GetStaffFromApi(c context.Context, comapnyId int) (any, error) {
	data := `[]`
	return &data, nil
}

func (s *Store) GetStoreStaff(c context.Context, taskId int) (*any, error) {
	return nil, nil
}
