package handler

import (
	"time"

	"github.com/D-building-anonymaizer/backend-service"
	"github.com/D-building-anonymaizer/backend-service/pkg/service"
	"github.com/gin-contrib/cors"
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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{h.services.GetUrl()},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}))

	api := router.Group("/")
	{
		api.POST("/api/exit", h.Exit)
		api.GET("/", h.Index)
	}
	analyze := router.Group("/analyze")
	{
		analyze.POST("/", h.FileReciever)
	}
	return router
}
