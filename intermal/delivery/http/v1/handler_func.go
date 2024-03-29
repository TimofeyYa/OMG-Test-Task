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
	TaskId int `form:"task_id"`
}

type getTaskStatusRes struct {
	Status string  `json:"status"`
	Uri    *string `json:"uri"`
}

func (h *Handler) getTaskStatus(c context.Context, data *getTaskStatusReq) (*getTaskStatusRes, *httpwrap.ErrorHTTP) {
	if data.TaskId == 0 {
		return nil, &httpwrap.ErrorHTTP{
			Code: 400,
			Msg:  "task_id is required",
		}
	}

	statusData, err := h.service.GetStatusTask(c, data.TaskId)
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
		uri := fmt.Sprintf("%s/v1/get_full_staff_list_load_data/%d", h.baseHost, data.TaskId)
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
	TaskId int `uri:"task_id"`
}

func (h *Handler) getStaff(c context.Context, data *getStaffReq) (*any, *httpwrap.ErrorHTTP) {
	if data.TaskId == 0 {
		return nil, &httpwrap.ErrorHTTP{
			Code: 400,
			Msg:  "task_id is required",
		}
	}

	staff, err := h.service.GetStaff(c, data.TaskId)
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

	return staff, nil
}
