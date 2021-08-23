package controller

import (
	"github.com/gin-gonic/gin"
	"mall/model"
	"mall/service"
	"net/http"
)


type OrderController struct {
	orderService service.OrderService
}

func (orderController *OrderController) Create(c *gin.Context)  {
	order := new(model.Order)

	err := c.ShouldBindJSON(&order)
	if err != nil {

	}

	orderRes, err := orderController.orderService.Create(order)

	c.JSON(http.StatusOK, gin.H{"msg": orderRes})
}

func (orderController *OrderController) Update(c *gin.Context)  {
	order := new(model.Order)

	err := c.ShouldBindJSON(&order)
	if err != nil {

	}

	res, err := orderController.orderService.Create(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": false})
	}

	c.JSON(http.StatusOK, gin.H{"msg": res})
}