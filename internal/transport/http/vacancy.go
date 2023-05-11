package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetVacancies(c *gin.Context) {
	Vacancies, err := h.service.Vacancy.GetVacancies()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, Vacancies)
}
