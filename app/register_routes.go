package app

import (
	"myapp/controller"
	"myapp/middleware"
	"myapp/repository"
	"myapp/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
    // Inisialisasi repository
    userRepository := repository.NewUserRepository(db)
    productRepository := repository.NewProductRepository(db)

    // Inisialisasi service
    authService := service.NewAuthService(userRepository)
    productService := service.NewProductService(productRepository)

    // Inisialisasi controller
    authController := controller.NewAuthController(authService)
    productController := controller.NewProductController(productService)

    // Auth routes
    r.POST("/login", authController.Login)
    r.POST("/register", authController.Register)

    // Group routes untuk produk
    productRoutes := r.Group("/products")
    productRoutes.Use(middleware.AuthMiddleware)
    {
        productRoutes.GET("/", productController.GetAllProducts)
        productRoutes.GET("/:id", productController.GetProductByID)
        productRoutes.POST("/", middleware.AdminMiddleware, productController.CreateProduct)
        productRoutes.PUT("/:id", middleware.AdminMiddleware, productController.UpdateProduct)
        productRoutes.DELETE("/:id", middleware.AdminMiddleware, productController.DeleteProduct)
    }
}
