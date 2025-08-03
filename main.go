package main

import (
	"os"

	"com.faisalkc/config"
	"com.faisalkc/models"
	"com.faisalkc/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Product{})
	config.DB.AutoMigrate(&models.Manufacturer{})
	config.DB.AutoMigrate(&models.Ingredient{})
	config.DB.AutoMigrate(&models.Additive{})
	config.DB.AutoMigrate(&models.ProductIngredient{})
	config.DB.AutoMigrate(&models.ProductAdditive{})

	r := gin.Default()
	routes.RegisterRoutes(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	r.Run(":" + port)
}
