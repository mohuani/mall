package service

import (
	"mall/model"
	"mall/repository"
)

type IProduct interface {
	Create(product model.Product) (*model.Product, error)
	Update(product model.Product) (bool, error)
	Delete(id int64) (bool, error)
	ExistById(id int64) bool
	GetById(id int64) (bool, error)
}

type ProductService struct {
	productRepository repository.ProductRepository
}

func (productService *ProductService) Create(product model.Product) (*model.Product, error){
	return productService.productRepository.Create(product)
}

func (productService *ProductService) Update(product model.Product) (bool, error) {
	return productService.productRepository.Update(product)
}

func (productService *ProductService) Delete(id int64) (bool, error) {
	return productService.productRepository.Delete(id)
}

func (productService *ProductService) ExistById(id int64) bool{
	return productService.productRepository.ExistById(id)
}

func (productService *ProductService) GetById(id int64) (*model.Product, error) {
	return productService.productRepository.GetById(id)
}
