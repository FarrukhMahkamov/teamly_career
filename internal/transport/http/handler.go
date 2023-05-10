package handler

import (
	"github.com/FarrukhMahkamov/teamly_career/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			JobCategory := v1.Group("/job-category")
			{
				JobCategory.GET("/", h.GetJobCategories)
			}
		}
	}

	return router
}
