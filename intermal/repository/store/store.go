package store

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"omg/intermal/models"
	"sync"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const InProgressStatus = "InProgress"
const DoneStatus = "Done"
const ErrStatus = "Error"

type Store struct {
	conn      *pgxpool.Pool
	apiUri    string
	tokenPair *models.TokenPair
}

func NewStore(conn *pgxpool.Pool, apiUri string, tokens *models.TokenPair) *Store {
	return &Store{
		conn:      conn,
		apiUri:    apiUri,
		tokenPair: tokens,
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

func (s *Store) GetTaskStatus(c context.Context, taskId int) (*models.Status, error) {
	sqlReq := `SELECT status FROM tasks WHERE id = $1`

	var status string
	if err := s.conn.QueryRow(c, sqlReq, taskId).Scan(&status); err != nil {
		if err.Error() == "no rows in result set" {
			return nil, models.ErrNoFound
		}
		return nil, err
	}

	return &models.Status{
		Name:   status,
		IsDone: status == DoneStatus,
	}, nil
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
	defer tx.Rollback(c)

	sqlSelecActive := `SELECT id, company_id FROM tasks WHERE status = $1 FOR UPDATE`

	// Получаем все активные таски и блокируем строки
	rows, err := tx.Query(c, sqlSelecActive, InProgressStatus)
	if err != nil {
		return fmt.Errorf("err get active tasks: %s", err)
	}

	wg := &sync.WaitGroup{}

	//  Делаем запросы к API
	//   TODO: Можно сделать ограничение запросов (Паттерн  WorkerPool)
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
	}

	// Закрываем канал
	go func() {
		wg.Wait()
		close(staffDataCh)
	}()

	//  Обрабатываем результат запросов
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
	}

	if err := tx.Commit(c); err != nil {
		return fmt.Errorf("err commit status: %s", err)
	}

	return nil
}

func (s *Store) GetStaffFromApi(c context.Context, comapnyId int) (any, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Timeout: 60 * time.Second, Transport: tr}

	url := fmt.Sprintf("%s/api/v1/company/%d/staff", s.apiUri, comapnyId)
	req, _ := http.NewRequest("GET", url, nil)

	authHeader := fmt.Sprintf("Bearer %s, User %s", s.tokenPair.Bearer, s.tokenPair.User)
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Accept", "application/vnd.yclients.v2+json")
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		thisErr := fmt.Errorf("err: request to api status code %d", res.StatusCode)
		log.Println(thisErr.Error())
		return nil, thisErr
	}

	var dataStaff any
	err = json.NewDecoder(res.Body).Decode(&dataStaff)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &dataStaff, nil
}

func (s *Store) GetStoreStaff(c context.Context, taskId int) (*any, error) {
	return nil, nil
}
