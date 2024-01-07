package v1

import (
	"context"
	httpwrap "omg/pkg/httpWrap"
)

type createTaskReq struct {
	CompanyId int `uri:"company_id" binding:"required"`
}

type createTaskRes struct {
	TaskId int `json:"task_id"`
}

func (h *Handler) createTask(r context.Context, data *createTaskReq) (*createTaskRes, *httpwrap.ErrorHTTP) {

	return &createTaskRes{
		TaskId: data.CompanyId,
	}, nil
}

type getTaskStatusReq struct {
	CompanyId int `uri:"company_id" binding:"required"`
	TaskIdId  int `uri:"task_id" binding:"required"`
}

type getTaskStatusRes struct {
	Status string  `json:"status"`
	Uri    *string `json:"uri"`
}

func (h *Handler) getTaskStatus(r context.Context, data *getTaskStatusReq) (*getTaskStatusRes, *httpwrap.ErrorHTTP) {

	return &getTaskStatusRes{}, nil
}
