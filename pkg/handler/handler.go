package handler

import (
	"github.com/D-building-anonymaizer/backend-service/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{h.services.GetUrl()},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}))

	api := router.Group("/api")
	{
		api.POST("/", h.function)
		api.GET("/", h.function)
	}
	analyze := router.Group("/analyze")
	{
		analyze.POST("/", h.fileReciever)
	}
	return router
}
