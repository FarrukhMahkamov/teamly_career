package handler

import (
	"log"

	"github.com/FarrukhMahkamov/teamly_career/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) FileUpload(c *gin.Context) {
	Destination := c.Param("destination")

	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	//there are 2 types of files: cv and photo
	FileName, err := pkg.UploadRequestFile(file, Destination)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"file": FileName,
	})
}

func (h *Handler) ServeFile(c *gin.Context) {
	FileName := c.Param("file_name")

	c.File("./assets/cv/" + FileName)
}
