package handler

import "github.com/gin-gonic/gin"

func (h *Handler) GetTeams(c *gin.Context) {
	Teams, err := h.service.Team.GetTeams()
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	c.JSON(200, Teams)
}
