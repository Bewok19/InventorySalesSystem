package entity

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string    `json:"username" validate:"required,min=3,max=20" gorm:"unique"`
	Email     string    `json:"email" validate:"required,email" gorm:"unique"`
	Password  string    `json:"password" validate:"required,min=6"`
	Role      string    `json:"role" gorm:"default:user"` // Default role is "user"
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"` // Automatically set on insert
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"` // Automatically set on update
}
