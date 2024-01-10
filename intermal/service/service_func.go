package service

import (
	"context"
	"log"
	"omg/intermal/models"
	"time"
)

func (s *service) CreateTask(c context.Context, companyId int) (int, error) {
	// TODO: Можно добавить проверку, есть ли уже активная таска для данной компании, если есть вернуть её id
	return s.repo.Store.CreateTask(c, companyId)
}

func (s *service) GetStatusTask(c context.Context, taskId int) (*models.Status, error) {
	return s.repo.Store.GetTaskStatus(c, taskId)
}

func (s *service) GetStaff(c context.Context, taskId int) (*any, error) {
	return s.repo.GetStoreStaff(c, taskId)
}

func (s *service) ResolveTasks(c context.Context) error {
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-c.Done():
			return nil
		case <-t.C:
			err := s.repo.PerformActiveTasks(c)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}
