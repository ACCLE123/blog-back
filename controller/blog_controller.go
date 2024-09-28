package controller

import (
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BlogController struct {
	BlogService *service.BlogService
}

func NewBlogController(blogService *service.BlogService) *BlogController {
	return &BlogController{
		BlogService: blogService,
	}
}

func (ctrl *BlogController) GetAllBlogs(c *gin.Context) {
	blogs, err := ctrl.BlogService.GetAllBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blogs)
}
