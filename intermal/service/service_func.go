package service

import "context"

func (s *service) CreateTask(c context.Context, companyId int) (int, error) {
	// TODO: Можно добавить проверку, есть ли уже активная таска для данной компании, если есть вернуть её id
	return s.repo.Store.CreateTask(c, companyId)
}

func (s *service) GetStatusTask(c context.Context, taskId int) (string, error) {
	return "", nil
}

func (s *service) GetStaff(c context.Context, taskId int) (*any, error) {
	return nil, nil
}

func (s *service) ResolveTasks(c context.Context) error {
	return nil
}
