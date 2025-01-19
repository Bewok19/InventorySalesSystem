package controller

import (
	"myapp/entity"
	"myapp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
    Register(ctx *gin.Context)
    Login(ctx *gin.Context)
}

type authController struct {
    authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
    return &authController{authService: authService}
}

func (c *authController) Register(ctx *gin.Context) {
    var request entity.RegisterRequest

    // Bind JSON dari request body ke struct RegisterRequest
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Panggil service untuk registrasi
    user, err := c.authService.Register(request)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Berikan respon sukses
    ctx.JSON(http.StatusOK, gin.H{
        "message": "Registration successful",
        "user": gin.H{
            "id":       user.ID,
            "username": user.Username,
            "email":    user.Email,
            "role":  user.Role,
        },
    })
}

func (c *authController) Login(ctx *gin.Context) {
    var request struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required"`
    }

    // Bind JSON dari request body ke struct
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Panggil service untuk proses login
    token, err := c.authService.Login(request.Email, request.Password)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    // Berikan respon sukses dengan token
    ctx.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "token":   token,
    })
}

