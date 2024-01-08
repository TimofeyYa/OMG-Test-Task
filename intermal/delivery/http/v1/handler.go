package v1

import (
	"omg/intermal/service"
	httpwrap "omg/pkg/httpWrap"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) BindRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/:company_id/staff/task", httpwrap.Wrap(h.createTask))
		v1.GET("/:company_id/staff/task/:task_id", httpwrap.Wrap(h.getTaskStatus))
		v1.GET("/:company_id/staff/:task_id", httpwrap.Wrap(h.getStaff))
	}
}
