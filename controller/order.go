package controller

import (
	"github.com/gin-gonic/gin"
	"mall/model"
	"mall/service"
	"net/http"
)

type OrderController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type orderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &orderController{
		orderService: orderService,
	}
}

func (orderController *orderController) Create(ctx *gin.Context) {
	var order model.Order

	err := ctx.ShouldBindJSON(&order)
	if err != nil {

	}

	res, err := orderController.orderService.Create(order)

	ctx.JSON(http.StatusOK, gin.H{"msg": res})
}

func (orderController *orderController) Update(ctx *gin.Context) {
	var order model.Order
	err := ctx.ShouldBindJSON(&order)
	if err != nil {

	}

	res, err := orderController.orderService.Update(order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": false})
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": res})
}
