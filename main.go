package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mall/config"
	"mall/controller"
	"mall/repository"
	"mall/service"
)

var (
	db                *gorm.DB                     = config.SetupDatabaseConnection()
	orderRepository   repository.OrderRepository   = repository.NewOrderRepository(db)
	productRepository repository.ProductRepository = repository.NewProductRepository(db)

	orderService   service.OrderService   = service.NewOrderService(orderRepository)
	productService service.ProductService = service.NewProductService(productRepository)

	orderController   controller.OrderController   = controller.NewOrderController(orderService)
	productController controller.ProductController = controller.NewProductController(productService)
)

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	order := router.Group("/order")
	{
		order.POST("/create", orderController.Create)
		order.POST("/update", orderController.Update)
	}

	product := router.Group("/product")
	{
		product.POST("/create", productController.Create)
		product.POST("/update", productController.Update)
	}

	router.Run()
}
