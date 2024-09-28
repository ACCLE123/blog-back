package main

import (
	"blog/controller"
	"blog/database"
	"blog/mapper"
	"blog/routes"
	"blog/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // 允许所有来源的请求
		// 也可以配置特定的来源
		// AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // 允许携带凭证（如 cookies）
	}))

	DB := database.Connect()
	blogMapper := mapper.NewBlogRepository(DB)
	blogService := service.NewBlogService(blogMapper)
	controller := controller.NewBlogController(blogService)

	routes.RegisterRoutes(r, controller)
	r.Run()
}
