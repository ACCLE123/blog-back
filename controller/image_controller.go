package controller

import (
	"blog/oss"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageController struct {
	OSS *oss.AliyunOSS
}

func NewImageController(oss *oss.AliyunOSS) *ImageController {
	return &ImageController{OSS: oss}
}

func (ic *ImageController) UploadImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file"})
		return
	}
	defer file.Close()

	url, err := ic.OSS.UploadImage(file, header.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": url})
}
