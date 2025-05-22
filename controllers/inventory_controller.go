package controllers

import (
	"inventory-system/config"
	"inventory-system/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInventory(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}