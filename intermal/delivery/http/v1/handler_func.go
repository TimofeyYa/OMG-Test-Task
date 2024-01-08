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

func (h *Handler) createTask(c context.Context, data *createTaskReq) (*createTaskRes, *httpwrap.ErrorHTTP) {
	newTaskId, err := h.service.CreateTask(c, data.CompanyId)
	if err != nil {
		return nil, &httpwrap.ErrorHTTP{
			Code: 500,
			Msg:  err.Error(),
		}
	}

	return &createTaskRes{
		TaskId: newTaskId,
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

type getStaffReq struct {
	CompanyId int `uri:"company_id" binding:"required"`
	TaskIdId  int `uri:"task_id" binding:"required"`
}

type getStaffRes struct {
	Data any `json:"data"`
}

func (h *Handler) getStaff(r context.Context, data *getStaffReq) (*getStaffRes, *httpwrap.ErrorHTTP) {

	return &getStaffRes{}, nil
}
