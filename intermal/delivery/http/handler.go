package http

import (
	v1 "omg/intermal/delivery/http/v1"
	"omg/intermal/service"

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

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	v1.NewHandler(h.service).BindRoutes(router)

	return router
}
