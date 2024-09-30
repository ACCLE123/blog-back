package controller

import (
	"blog/model"
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BlogController struct {
	BlogService *service.BlogService
}

func NewBlogController(blogService *service.BlogService) *BlogController {
	return &BlogController{
		BlogService: blogService,
	}
}

func (ctrl *BlogController) GetAllBlogs(ctx *gin.Context) {
	blogs, err := ctrl.BlogService.GetAllBlogs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, blogs)
}

func (ctrl *BlogController) GetBlogByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	blog, err := ctrl.BlogService.GetBlogByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

func (c *BlogController) AddBlog(ctx *gin.Context) {
	var blog model.Blog
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := c.BlogService.AddBlog(&blog); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog"})
		return
	}

	ctx.JSON(http.StatusOK, "Blog created successfully")
}
