package controller

import (
	"github.com/gin-gonic/gin"
	"mall/model"
	"mall/service"
	"net/http"
)

type ProductController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
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

	productRes, err := productController.productService.Create(product)

	ctx.JSON(http.StatusOK, gin.H{"msg": productRes})
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
