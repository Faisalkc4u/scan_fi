package routes

import (
	"com.faisalkc/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	productRoutes := r.Group("/products")
	{
		productRoutes.GET("", controllers.GetProducts)
		productRoutes.POST("", controllers.CreateProduct)
		productRoutes.GET("/:id", controllers.GetProductByID)
		productRoutes.PUT("/:id", controllers.UpdateProduct)
		productRoutes.DELETE("/:id", controllers.DeleteProduct)
		productRoutes.GET("/search/:barcode", controllers.SearchProductByBarcode)
	}
}
