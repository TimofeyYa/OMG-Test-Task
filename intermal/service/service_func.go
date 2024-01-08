package service

import "context"

func (s *service) CreateTask(c context.Context, companyId int) (int, error) {
	return 0, nil
}

func (s *service) GetStatusTask(c context.Context, taskId int) (string, error) {
	return "", nil
}

func (s *service) GetStaff(c context.Context, taskId int) (*any, error) {
	return nil, nil
}
