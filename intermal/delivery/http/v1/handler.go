package v1

import (
	httpwrap "omg/pkg/httpWrap"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) BindRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/:company_id/staff/task", httpwrap.Wrap(h.createTask))
		v1.GET("/:company_id/staff/task/:task_id", httpwrap.Wrap(h.getTaskStatus))
		v1.GET("/:company_id/staff/:task_id", httpwrap.Wrap(h.getStaff))
	}
}
