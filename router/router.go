package router

import (
	"myapp/controller"
	"myapp/middleware"
	"myapp/repository"
	"myapp/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    // Buat instance repository dan service
    productRepository := repository.NewProductRepository(db)
    productService := service.NewProductService(productRepository)
    productController := controller.NewProductController(productService)

    // Group routes untuk produk
    productRoutes := r.Group("/products")
    productRoutes.Use(middleware.AuthMiddleware)
    {
        productRoutes.GET("", productController.GetAllProducts)
        productRoutes.GET("/:id", productController.GetProductByID)
        productRoutes.POST("", middleware.AdminMiddleware, productController.CreateProduct)
        productRoutes.PUT("/:id", middleware.AdminMiddleware, productController.UpdateProduct)
        productRoutes.DELETE("/:id", middleware.AdminMiddleware, productController.DeleteProduct)
    }

    return r
}
