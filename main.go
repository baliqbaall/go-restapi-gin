package main

import (
	"github.com/baliqbaall/go-restapi-gin/controllers"
	"github.com/baliqbaall/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDB()

	r.GET("api/products", controllers.Index)
	r.GET("api/product/:id", controllers.Show)
	r.POST("api/product", controllers.Create)
	r.PUT("api/product/:id", controllers.Update)
	r.DELETE("api/product", controllers.Delete)

	r.Run()
}
