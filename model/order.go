package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId    int64
	ProductId int64
	Num       int64
	Money     float64
}
