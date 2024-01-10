package v1

import (
	"omg/intermal/service"
	httpwrap "omg/pkg/httpWrap"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service  service.Service
	baseHost string
}

func NewHandler(service service.Service, baseHost string) *Handler {
	return &Handler{
		service:  service,
		baseHost: baseHost,
	}
}

func (h *Handler) BindRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/get_full_staff_list", httpwrap.Wrap(h.createTask))
		v1.GET("/get_task_status", httpwrap.Wrap(h.getTaskStatus))
		v1.GET("/get_full_staff_list_load_data/:task_id", httpwrap.Wrap(h.getStaff))
	}
}
