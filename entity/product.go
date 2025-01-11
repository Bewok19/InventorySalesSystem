package entity

type Product struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
	Price int    `gorm:"type:int;not null" json:"price" binding:"required,gt=0"`
	Stock int    `gorm:"type:int;not null" json:"stock" binding:"gte=0"`
}

