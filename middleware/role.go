package middleware

import (
	"log"
	"myapp/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RoleMiddleware adalah middleware untuk memeriksa peran pengguna
func RoleMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil peran pengguna dari context (disimpan setelah verifikasi JWT)
		userRole, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Role not found"})
			c.Abort()
			return
		}

		// Konversi userRole ke string
		role := userRole.(string)

		// Periksa apakah peran pengguna ada di daftar peran yang diizinkan
		for _, r := range requiredRoles {
			if role == r {
				c.Next()
				return
			}
		}

		// Jika peran tidak diizinkan, kembalikan 403 Forbidden
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		c.Abort()
	}
}

// AuthMiddleware verifies the user's token.
func AuthMiddleware(c *gin.Context) {
    // Ambil token dari header Authorization
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
        c.Abort()
        return
    }

    token := strings.TrimPrefix(authHeader, "Bearer ")

    // Validasi token
    claims, err := utils.ValidateToken(token)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        c.Abort()
        return
    }

    // Set informasi user ke context
    c.Set("user_id", claims["user_id"])
    c.Set("role", claims["role"])
    c.Next()
}

// AdminMiddleware ensures the user has admin privileges.
func AdminMiddleware(c *gin.Context) {
    role, exists := c.Get("role")
    if !exists {
        log.Println("[DEBUG] Role not found in context")
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        c.Abort()
        return
    }

    log.Printf("[DEBUG] User role: %v", role)

    if role != "admin" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        c.Abort()
        return
    }

    c.Next()
}

