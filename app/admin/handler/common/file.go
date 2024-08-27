package common

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/utils"
)

type File struct{}

func NewFileRoute() *File {
	return &File{}
}

func (this *File) FileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}

	// Upload the file to your server.
	dst := "./static/upload/" + file.Filename
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}

	utils.OkWithData(c, gin.H{
		"message": "ok",
		"url":     "/upload/" + file.Filename,
	})
	return
}
