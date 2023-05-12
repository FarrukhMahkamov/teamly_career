package handler

import (
	"strconv"

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

func (h *Handler) UserPersonalInfoUpdate(c *gin.Context) {
	UserID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	var UserPersonalInfoUpdateRequest core.UserPersonalInfoUpdateRequest
	if err := c.BindJSON(&UserPersonalInfoUpdateRequest); err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	if err := h.service.User.UserPersonalInfoUpdate(UserPersonalInfoUpdateRequest, int64(UserID)); err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"Message": "Profile details updated successfully",
	})
}

func (h *Handler) UpdateUserPassword(c *gin.Context) {
	UserID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	var UpdateUserPasswordRequest core.UpdateUserPasswordRequest
	if err := c.BindJSON(&UpdateUserPasswordRequest); err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	if err := h.service.User.UpdateUserPassword(int64(UserID), UpdateUserPasswordRequest); err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"Message": "Password updated successfully",
	})
}
