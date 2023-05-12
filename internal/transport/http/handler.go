package handler

import (
	"github.com/FarrukhMahkamov/teamly_career/internal/service"
	"github.com/FarrukhMahkamov/teamly_career/internal/transport/http/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/sign-up", h.SignUp)
	router.POST("/sign-in", h.SignIn)

	router.POST("/file-upload/:destination", h.FileUpload)
	router.GET("/serve-file/:destination/:file_name", h.ServeFile)

	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/teams", h.GetTeams)

			vacancies := v1.Group("/vacancies")
			{
				vacancies.GET("/", h.GetVacancies)
				vacancies.POST("/", h.AddVacancy)
			}

			users := v1.Group("/users")
			{
				users.PUT("/personal-info/:user_id", h.UserPersonalInfoUpdate)
				users.PUT("/password/:user_id", h.UpdateUserPassword)
			}
		}
	}

	return router
}
