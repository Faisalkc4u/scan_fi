package main

import (
	"com.faisalkc/config"
	"com.faisalkc/models"
	"com.faisalkc/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Product{})

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8000")
}
