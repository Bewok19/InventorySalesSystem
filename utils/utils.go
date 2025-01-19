package utils

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Secret key untuk JWT (pastikan disimpan aman, misalnya di env)
var jwtSecretKey = []byte("your_secret_key")

// ValidateToken memvalidasi JWT token
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return jwtSecretKey, nil
    })

    if err != nil || !token.Valid {
        return nil, errors.New("invalid token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return nil, errors.New("invalid token claims")
    }

    return claims, nil
}

// ExtractClaimsFromContext mengambil klaim dari context request
func ExtractClaimsFromContext(c *gin.Context) (map[string]interface{}, error) {
	claims, exists := c.Get("claims") // "claims" disimpan sebelumnya di middleware JWT
	if !exists {
		return nil, errors.New("no claims found in context")
	}

	if claimsMap, ok := claims.(map[string]interface{}); ok {
		return claimsMap, nil
	}

	return nil, errors.New("invalid claims format")
}

// HashPassword membuat hash dari password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword membandingkan hash password dengan password plain
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// VerifyPassword membandingkan hash password dengan plain password
func VerifyPassword(hashedPassword, plainPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
    return err == nil
}

// GenerateToken membuat JWT token
func GenerateToken(userID uint, role string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "role":    role,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecretKey) // Same key used in AuthMiddleware
}
