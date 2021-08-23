package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName  string `json:"product_name"`
	ProductImage string `json:"product_image"`
	ProductPrice float64 `json:"product_price"`
}
