package handler

import (
	"github.com/D-building-anonymaizer/backend-service"
	"github.com/D-building-anonymaizer/backend-service/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	server   *backend.Server
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Static("/static", "../../build/static")
	router.POST("/api/exit", h.Exit)
	router.GET("/", h.Index)
	//router.POST("/mail", h.MailSender)
	router.POST("/analize", h.FileReciever)

	return router
}
