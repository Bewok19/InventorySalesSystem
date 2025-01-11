package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
    r.GET("/products", GetAllProducts)
    r.POST("/products", CreateProduct)
}
