package routes

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, blogController *controller.BlogController, imageController *controller.ImageController) {
	r.GET("/ping", controller.Ping)
	r.GET("/blogs", blogController.GetAllBlogs)
	r.GET("/blogs/:id", blogController.GetBlogByID)
	r.POST("/blogs", blogController.AddBlog)
	r.PUT("/blogs", blogController.UpdateBlog)
	r.DELETE("/blogs/:id", blogController.DeleteBlog)
	r.POST("/blogs/updateOrAdd", blogController.UpdateOrAddBlog)
	r.POST("/upload/image", imageController.UploadImage)
}
