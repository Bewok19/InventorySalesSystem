package repository

import (
	"errors"
	"fmt"
	"myapp/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*entity.User, error)
	Save(user *entity.User) error
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
    var user entity.User
    if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            fmt.Println("User not found for email:", email) // Logging untuk debugging
            return nil, nil
        }
        fmt.Println("Error finding user by email:", err) // Logging error
        return nil, err
    }
    fmt.Println("User found:", user) // Logging sukses
    return &user, nil
}


func (r *userRepositoryImpl) Save(user *entity.User) error {
    fmt.Println("Saving user:", user) // Logging sebelum menyimpan
    if err := r.db.Create(user).Error; err != nil {
        fmt.Println("Error saving user:", err) // Logging error
        return err
    }
    fmt.Println("User saved successfully:", user) // Logging sukses
    return nil
}