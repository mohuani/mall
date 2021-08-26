package service

import (
	"mall/model"
	"mall/repository"
)

type ProductService interface {
	Create(product model.Product) (*model.Product, error)
	Update(product model.Product) (bool, error)
	Delete(product model.Product) (bool, error)
	ExistById(id int64) bool
	GetById(id int64) (*model.Product, error)
	List() ([]*model.Product, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepository,
	}
}

func (productService *productService) Create(product model.Product) (*model.Product, error) {
	return productService.productRepository.Create(product)
}

func (productService *productService) Update(product model.Product) (bool, error) {
	return productService.productRepository.Update(product)
}

func (productService *productService) Delete(product model.Product) (bool, error) {
	return productService.productRepository.Delete(product)
}

func (productService *productService) ExistById(id int64) bool {
	return productService.productRepository.ExistById(id)
}

func (productService *productService) GetById(id int64) (*model.Product, error) {
	return productService.productRepository.GetById(id)
}

func (productService *productService) List() ([]*model.Product, error) {
	return productService.productRepository.List()
}
