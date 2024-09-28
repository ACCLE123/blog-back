package main

import (
	"blog/controller"
	"blog/database"
	"blog/mapper"
	"blog/routes"
	"blog/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	DB := database.Connect()
	blogMapper := mapper.NewBlogRepository(DB)
	blogService := service.NewBlogService(blogMapper)
	controller := controller.NewBlogController(blogService)

	routes.RegisterRoutes(r, controller)
	r.Run()
}
