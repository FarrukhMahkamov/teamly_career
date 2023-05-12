package handler

import (
	"log"

	"github.com/FarrukhMahkamov/teamly_career/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) FileUpload(c *gin.Context) {
	Destination := c.Param("destination")

	File, _ := c.FormFile("file")
	log.Println(File.Filename)

	//there are 2 types of files: cv and photo
	FileName, err := pkg.UploadRequestFile(File, Destination)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"uploaded_file": FileName,
	})
}

func (h *Handler) ServeFile(c *gin.Context) {
	FileName := c.Param("file_name")
	Destination := c.Param("destination")

	c.File("./assets/images/" + Destination + "/" + FileName)
}
