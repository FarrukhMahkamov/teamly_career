package handler

import (
	"net/http"

	"github.com/FarrukhMahkamov/teamly_career/internal/core"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetVacancies(c *gin.Context) {
	Vacancies, err := h.service.Vacancy.GetVacancies()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, Vacancies)
}

func (h *Handler) AddVacancy(c *gin.Context) {
	var AddVacancyRequest core.AddVacancyRequest

	if err := c.BindJSON(&AddVacancyRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	VacancyRequest := core.VacancyRequest{
		VacancyTeamId:      AddVacancyRequest.VacancyTeamId,
		VacancyStatusId:    AddVacancyRequest.VacancyStatusId,
		Position:           AddVacancyRequest.Position,
		VacancyDescription: AddVacancyRequest.VacancyDescription,
	}

	VacancyId, err := h.service.Vacancy.AddVacancy(VacancyRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}

	// VacancyDetailRequest := core.VacancyDetailRequest{
	// 	VacancyId:    VacancyId,
	// 	Level:        AddVacancyRequest.Level,
	// 	Experience:   AddVacancyRequest.Experience,
	// 	WorkType:     AddVacancyRequest.WorkType,
	// 	WorkTime:     AddVacancyRequest.WorkTime,
	// 	WorkLocation: AddVacancyRequest.WorkLocation,
	// 	Salary:       AddVacancyRequest.Salary,
	// }

	// _, err = h.service.VacancyDetail.AddVacancyDetail(VacancyDetailRequest)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	// }

	c.JSON(http.StatusOK, VacancyId)
}
