package controllers

import (
	"inventory-system/config"
	"inventory-system/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePurchaseOrder(c *gin.Context) {
	var order models.PurchaseOrder
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	if err := config.DB.First(&product, order.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	product.Quantity += order.Quantity
	config.DB.Save(&product)
	config.DB.Create(&order)
	c.JSON(http.StatusOK, order)
}