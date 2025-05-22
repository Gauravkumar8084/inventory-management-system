package controllers

import (
	"inventory-system/config"
	"inventory-system/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSalesOrder(c *gin.Context) {
	var order models.SalesOrder
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	if err := config.DB.First(&product, order.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if product.Quantity < order.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
		return
	}

	product.Quantity -= order.Quantity
	order.TotalPrice = float64(order.Quantity) * product.Price

	config.DB.Save(&product)
	config.DB.Create(&order)
	c.JSON(http.StatusOK, order)
}