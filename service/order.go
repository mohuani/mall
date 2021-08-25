package service

import (
	"mall/model"
	"mall/repository"
)

type OrderService interface {
	Create(order model.Order) (*model.Order, error)
	Update(order model.Order) (bool, error)
	Delete(id int64) (bool, error)
	ExistById(id int64) bool
	GetById(id int64) (*model.Order, error)
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderService{
		orderRepository: orderRepository,
	}
}

func (orderService *orderService) Create(order model.Order) (*model.Order, error) {
	return orderService.orderRepository.Create(order)
}

func (orderService *orderService) Update(order model.Order) (bool, error) {
	return orderService.orderRepository.Update(order)
}

func (orderService *orderService) Delete(id int64) (bool, error) {
	return orderService.orderRepository.Delete(id)
}

func (orderService *orderService) ExistById(id int64) bool {
	return orderService.orderRepository.ExistById(id)
}

func (orderService *orderService) GetById(id int64) (*model.Order, error) {
	return orderService.orderRepository.GetById(id)
}
