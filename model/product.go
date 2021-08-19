package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName  string
	ProductImage string
	ProductPrice float64
}
