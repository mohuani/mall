package service

import (
	"mall/model"
	"mall/repository"
)

type IOrder interface {
	Create(order model.Order) (*model.Order, error)
	Update(order model.Order) (bool, error)
	Delete(id int64) (bool, error)
	ExistById(id int64) bool
	GetById(id int64) (*model.Order, error)
}

type OrderService struct {
	orderRepository repository.OrderRepository
}

func (orderService *OrderService) Create(order *model.Order) (*model.Order, error) {
	return orderService.orderRepository.Create(order)
}

func (orderService *OrderService) Update(order model.Order) (bool, error) {
	return orderService.orderRepository.Update(order)
}

func (orderService *OrderService) Delete(id int64) (bool, error) {
	return orderService.orderRepository.Delete(id)
}

func (orderService *OrderService) ExistById(id int64) bool{
	return orderService.orderRepository.ExistById(id)
}

func (orderService *OrderService) GetById(id int64) (*model.Order, error) {
	return orderService.orderRepository.GetById(id)
}