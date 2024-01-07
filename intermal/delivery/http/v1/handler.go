package v1

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) BindRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/:company_id/staff/task")
		v1.GET("/:company_id/staff/task/:task_id")
		v1.GET("/:company_id/staff/:task_id")
	}
}
