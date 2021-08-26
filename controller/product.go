package controller

import (
	"github.com/gin-gonic/gin"
	"mall/model"
	"mall/service"
	"net/http"
	"strconv"
)

type ProductController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	List(ctx *gin.Context)
}

type productController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}

func (productController *productController) Create(ctx *gin.Context) {
	var product model.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {

	}

	res, err := productController.productService.Create(product)

	ctx.JSON(http.StatusOK, gin.H{"msg": res})
}

func (productController *productController) Update(ctx *gin.Context) {
	var product model.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {

	}

	res, err := productController.productService.Update(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": false})
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": res})
}

func (productController *productController) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Query("id"), 0, 0)
	var product model.Product

	product.ID = uint(id)

	err := ctx.ShouldBindJSON(&product)
	if err != nil {

	}

	res, err := productController.productService.Delete(product)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": res})

}

func (productController *productController) List(ctx *gin.Context) {
	res, err := productController.productService.List()
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": res})

}
