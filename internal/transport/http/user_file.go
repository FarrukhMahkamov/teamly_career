package handler

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/FarrukhMahkamov/teamly_career/internal/core"
	"github.com/FarrukhMahkamov/teamly_career/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddUserFile(c *gin.Context) {
	File, err := c.FormFile("file")
	if err != nil {
		log.Println("Error reading form file:", err)
		c.JSON(400, gin.H{
			"error": "file not found in request",
		})
		return
	}

	Destination := "cv"

	UploadFiled, err := pkg.UploadRequestFile(File, Destination)
	if err != nil {
		log.Println("Error uploading file:", err)
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	UserFileRequest := core.UserFileRequest{
		UserID:   int(c.MustGet("user_id").(int64)),
		FileName: UploadFiled,
		FileSize: File.Size,
		FileType: filepath.Ext(File.Filename),
		StatusID: 1,
	}

	err = h.service.AddUserFile(UserFileRequest)
	if err != nil {
		log.Println("Error adding user file:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"message": "File uploaded successfully",
	})
}
