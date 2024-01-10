package http

import (
	v1 "omg/intermal/delivery/http/v1"
	"omg/intermal/service"

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

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	v1.NewHandler(h.service, h.baseHost).BindRoutes(router)

	return router
}
