package http

import (
	v1 "omg/intermal/delivery/http/v1"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	v1.NewHandler().BindRoutes(router)

	return router
}
