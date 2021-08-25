package repository

import (
	"gorm.io/gorm"
	"log"
	"mall/model"
)

type OrderRepository interface {
	Create(order model.Order) (*model.Order, error)

	Update(order model.Order) (bool, error)

	GetById(id int64) (*model.Order, error)

	Delete(id int64) (bool, error)

	ExistById(id int64) bool
}

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		DB: db,
	}
}

func (orderRepository *orderRepository) Create(order model.Order) (*model.Order, error) {
	err := orderRepository.DB.Create(order).Error
	if err != nil {
		log.Println("order创建失败")
		return &order, err
	}

	return &order, nil
}

func (orderRepository *orderRepository) Update(order model.Order) (bool, error) {
	originOrder := new(model.Order)

	err := orderRepository.DB.Where("id=?", order.ID).First(originOrder).Updates(map[string]interface{}{
		"user_id":    order.UserId,
		"product_id": order.ProductId,
		"num":        order.Num,
		"money":      order.Money,
	}).Error

	if err != nil {
		log.Println("order更新失败")
		return false, err
	}

	return true, nil
}

func (orderRepository *orderRepository) GetById(id int64) (*model.Order, error) {
	order := new(model.Order)

	err := orderRepository.DB.Where("id=?", id).First(order).Error
	if err != nil {
		log.Println("order查询失败")
		return nil, err
	}

	return order, nil
}

func (orderRepository *orderRepository) Delete(id int64) (bool, error) {
	err := orderRepository.DB.Where("id=?", id).Delete(model.Order{}).Error
	if err != nil {
		log.Println("order删除失败")
		return false, err
	}

	return true, nil
}

func (orderRepository *orderRepository) ExistById(id int64) bool {
	err := orderRepository.DB.Where("id=?", id).First(model.Order{}).Error
	if err != nil {
		return false
	}

	return true
}
