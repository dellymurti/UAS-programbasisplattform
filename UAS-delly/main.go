package main

import (
	"UAS-DELLY/produk/produkhandler"
	"UAS-DELLY/produk/userhandler"
	"UAS-DELLY/models"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "UAS-DELLY/docs"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	docs.SwaggerInfo.BasePath = "/api"

	api := r.Group("/api")
	{
				api.POST("/login", userhandler.Login)
				produk := api.Group("/produk")
		{
			produk.GET("/", produkhandler.Index)
			produk.GET("/:id", produkhandler.Show)
			produk.POST("/", produkhandler.Create)
			produk.PUT("/:id", produkhandler.Update)
			produk.DELETE("/:id", produkhandler.Delete)
		}
			user := api.Group("/user")
		{
			user.GET("/", userhandler.Index)
			user.GET("/:id", userhandler.Show)
			user.POST("/", userhandler.Create)
			user.PUT("/:id", userhandler.Update)
			user.DELETE("/:id", userhandler.Delete)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
