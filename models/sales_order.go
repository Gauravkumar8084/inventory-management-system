package models

import "gorm.io/gorm"

type SalesOrder struct {
	gorm.Model
	ProductID  uint    `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	OrderDate  string  `json:"order_date"`
}