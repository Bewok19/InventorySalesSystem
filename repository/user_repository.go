package repository

import (
	"myapp/config"
	"myapp/entity"
)

func ValidateUser(username, password string) bool {
    var user entity.User
    result := config.DB.Where("username = ? AND password = ?", username, password).First(&user)
    return result.Error == nil
}
