package router

import (
	"myapp/controller"
	"myapp/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Rute login
    r.POST("/login", controller.Login)

    // Rute produk
    productRoutes := r.Group("/products", middleware.JWTAuth())
    {
        productRoutes.GET("/:id", controller.GetProductByID)
        productRoutes.POST("/", controller.CreateProduct)
        productRoutes.PUT("/:id", controller.UpdateProduct)
        productRoutes.DELETE("/:id", controller.DeleteProduct)
    }

    return r
}
