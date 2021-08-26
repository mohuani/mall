package repository

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"mall/model"
)

type ProductRepository interface {
	Create(product model.Product) (*model.Product, error)
	Update(product model.Product) (bool, error)
	Delete(product model.Product) (bool, error)
	ExistById(id int64) bool
	GetById(id int64) (*model.Product, error)
	List() ([]*model.Product, error)
}

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		DB: db,
	}
}

func (productRepository *productRepository) Create(product model.Product) (*model.Product, error) {
	err := productRepository.DB.Create(&product).Error
	if err != nil {
		log.Println("product创建失败")
		return &product, err
	}

	return &product, nil
}

func (productRepository *productRepository) Update(product model.Product) (bool, error) {
	originProduct := new(model.Product)

	err := productRepository.DB.Where("id=?", product.ID).First(originProduct).Updates(map[string]interface{}{
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

func (productRepository *productRepository) Delete(product model.Product) (bool, error) {
	fmt.Println(product.ID)
	err := productRepository.DB.Where("id=?", product.ID).Delete(&model.Product{}).Error
	if err != nil {
		log.Println("product删除失败")
		return false, err
	}

	return true, nil
}

func (productRepository *productRepository) ExistById(id int64) bool {
	err := productRepository.DB.Where("id=?", id).First(model.Product{}).Error
	if err != nil {
		return false
	}

	return true
}

func (productRepository *productRepository) GetById(id int64) (*model.Product, error) {
	product := new(model.Product)

	err := productRepository.DB.Where("id=?", id).First(product).Error
	if err != nil {
		log.Println("product查询失败")
		return nil, err
	}

	return product, nil
}

func (productRepository *productRepository) List() (products []*model.Product, err error) {
	err = productRepository.DB.Find(&products).Error
	if err != nil {
		log.Println("product查询失败")
		return nil, err
	}

	return products, nil
}
