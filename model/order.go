package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId    int64   `json:"user_id"`
	ProductId int64   `json:"product_id"`
	Num       int64   `json:"num"`
	Money     float64 `json:"money"`
}
