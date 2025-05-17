package controllers

import (
	"errors"
	"net/http"

	"com.faisalkc/config"
	"com.faisalkc/models"
	"com.faisalkc/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	var product models.Product

	// Parse JSON
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if product.Barcode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Barcode is required"})
		return
	}
	// Check if barcode already exists
	var existing models.Product
	err := config.DB.Where("barcode = ?", product.Barcode).First(&existing).Error

	if err == nil {
		// Found — conflict
		c.JSON(http.StatusConflict, gin.H{"error": "Product with this barcode already exists"})
		return
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Unexpected DB error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Create new product
	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to create product",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	config.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
func SearchProductByBarcode(c *gin.Context) {
	id := c.Param("barcode")
	var product models.Product

	// Try to find existing product
	err := config.DB.
		Preload("Brand").
		Preload("Ingredients").
		Preload("Additives").
		Where("barcode = ?", id).First(&product).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create new placeholder product
			product = models.Product{
				Barcode:        id,
				NumberOfSearch: 1,
			}
			if err := config.DB.Create(&product).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create new product"})
				utils.LogError("Failed to create placeholder product", err)
				return
			}
			utils.LogInfo("Placeholder product created", " Barcode: "+id)
			c.JSON(http.StatusOK, product)
			return
		}

		// Other DB error
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		utils.LogError("Database error", err)
		return
	}

	// Product found — increment search count
	product.NumberOfSearch++
	config.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}
