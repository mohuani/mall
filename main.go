package main

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"mall/config"
	"mall/controller"
	"mall/repository"
	"mall/service"
	"net/http"
)

var (
	db                *gorm.DB                     = config.SetupDatabaseConnection()
	orderRepository   repository.OrderRepository   = repository.NewOrderRepository(db)
	productRepository repository.ProductRepository = repository.NewProductRepository(db)
	//esRepository repository.EsRepository = repository.NewEsRepository(db)

	orderService   service.OrderService   = service.NewOrderService(orderRepository)
	productService service.ProductService = service.NewProductService(productRepository)
	//esService service.EsService = service.NewEsService(esRepository)

	orderController   controller.OrderController   = controller.NewOrderController(orderService)
	productController controller.ProductController = controller.NewProductController(productService)
	//esController controller.EsController = controller.NewEsController(esService)
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
		product.GET("/delete", productController.Delete)
		product.GET("/list", productController.List)
	}

	es := router.Group("/es")
	{
		es.GET("/create", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"msg": "1234"})

		})

		es.GET("/get", func(context *gin.Context) {

			es, err := elasticsearch.NewDefaultClient()
			if err != nil {
				return
			}
			log.Println(elasticsearch.Version)
			log.Println(es.Info())

			context.JSON(http.StatusOK, gin.H{"msg": elasticsearch.Version})

		})

	}

	router.Run()
}
