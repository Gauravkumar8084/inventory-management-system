package routes

import (
	"github.com/gin-gonic/gin"
	"inventory-system/controllers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/api/products", controllers.GetProducts)
	r.POST("/api/products", controllers.CreateProduct)
	r.GET("/api/products/:id", controllers.GetProduct)
	r.PUT("/api/products/:id", controllers.UpdateProduct)
	r.DELETE("/api/products/:id", controllers.DeleteProduct)

	r.POST("/api/sales-orders", controllers.CreateSalesOrder)
	r.POST("/api/purchase-orders", controllers.CreatePurchaseOrder)
	r.GET("/api/inventory", controllers.GetInventory)

	return r
}