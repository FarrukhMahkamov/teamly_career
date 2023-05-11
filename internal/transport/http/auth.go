package handler

import (
	"github.com/FarrukhMahkamov/teamly_career/internal/core"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var UserRequest core.UserRequest
	if err := c.BindJSON(&UserRequest); err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	UserID, err := h.service.User.RegisterUser(UserRequest)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"User_id": UserID,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	var UserLoginRequest core.UserLoginRequest
	if err := c.BindJSON(&UserLoginRequest); err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	LoginedUser, err := h.service.Auth.LoginUser(UserLoginRequest)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	c.JSON(200, LoginedUser)
}
