package repository

import (
	"gorm.io/gorm"
	"log"
	"mall/model"
)

type IOrderRepository interface {
	Create(product model.Order) (*model.Order, error)

	Update(product model.Order) (bool, error)

	GetById(id int64) (*model.Order, error)

	Delete(id int64) (bool, error)

	ExistById(id int64) bool
}

type OrderRepository struct {
	DB gorm.DB
}

func (OrderRepository *OrderRepository) Create(product model.Product) (*model.Product, error) {
	err := OrderRepository.DB.Create(product).Error
	if err != nil {
		log.Println("product创建失败")
		return &product, err
	}

	return &product, nil
}

func (OrderRepository *OrderRepository) Update(product model.Product) (bool, error) {
	originProduct := new(model.Product)

	err := OrderRepository.DB.Where("id=?", product.ID).First(originProduct).Updates(map[string]interface{}{
		"product_name":  product.ProductName,
		"product_image": product.ProductImage,
		"product_price": product.ProductPrice,
	}).Error

	if err != nil {
		log.Println("product更新失败")
		return false, err
	}

	return true, nil
}

func (OrderRepository *OrderRepository) GetById(id int64) (*model.Product, error) {
	product := new(model.Product)

	err := OrderRepository.DB.Where("id=?", id).First(product).Error
	if err != nil {
		log.Println("product查询失败")
		return nil, err
	}

	return product, nil
}

func (OrderRepository *OrderRepository) Delete(id int64) (bool, error) {
	err := OrderRepository.DB.Where("id=?", id).Delete(model.Product{}).Error
	if err != nil {
		log.Println("product删除失败")
		return false, err
	}

	return true, nil
}

func (OrderRepository *OrderRepository) ExistById(id int) bool {
	err := OrderRepository.DB.Where("id=?", id).First(model.Product{}).Error
	if err != nil {
		return false
	}

	return true
}
