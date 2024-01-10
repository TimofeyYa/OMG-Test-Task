package v1

import (
	"context"
	"errors"
	"fmt"
	"omg/intermal/models"
	httpwrap "omg/pkg/httpWrap"
)

type createTaskReq struct {
	CompanyId int `form:"company_id"`
}

type createTaskRes struct {
	TaskId int `json:"task_id"`
}

func (h *Handler) createTask(c context.Context, data *createTaskReq) (*createTaskRes, *httpwrap.ErrorHTTP) {
	if data.CompanyId == 0 {
		return nil, &httpwrap.ErrorHTTP{
			Code: 400,
			Msg:  "company_id is required",
		}
	}

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

func (h *Handler) getTaskStatus(c context.Context, data *getTaskStatusReq) (*getTaskStatusRes, *httpwrap.ErrorHTTP) {
	statusData, err := h.service.GetStatusTask(c, data.CompanyId, data.TaskIdId)
	if err != nil {
		if errors.Is(models.ErrNoFound, err) {
			return nil, &httpwrap.ErrorHTTP{
				Code: 404,
				Msg:  err.Error(),
			}
		}
		return nil, &httpwrap.ErrorHTTP{
			Code: 500,
			Msg:  err.Error(),
		}
	}

	if statusData.IsDone {
		uri := fmt.Sprintf("/%d/staff/%d", data.CompanyId, data.TaskIdId)
		return &getTaskStatusRes{
			Uri:    &uri,
			Status: statusData.Name,
		}, nil
	}

	return &getTaskStatusRes{
		Uri:    nil,
		Status: statusData.Name,
	}, nil
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
