package entity

import (
	"time"
)

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" validate:"required,min=3,max=100"`
	Price       float64   `json:"price" validate:"required,gte=0"`
	Stock       int       `json:"stock" validate:"required,gte=0"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"` // Automatically set on insert
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"` // Automatically set on update
}
