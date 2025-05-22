package models

import "gorm.io/gorm"

type PurchaseOrder struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Supplier  string `json:"supplier"`
	OrderDate string `json:"order_date"`
}

