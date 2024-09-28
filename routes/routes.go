package routes

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, blogController *controller.BlogController) {
	r.GET("/ping", controller.Ping)
	r.GET("/blogs", blogController.GetAllBlogs)
	r.GET("/blogs/:id", blogController.GetBlogByID)
}
