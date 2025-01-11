package repository

import (
	"errors"
	"myapp/config"
	"myapp/entity"

	"gorm.io/gorm"
)

func GetAllProducts() ([]entity.Product, error) {
	var products []entity.Product
	result := config.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func GetProductByID(id uint) (entity.Product, error) {
	var product entity.Product
	result := config.DB.First(&product, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return product, errors.New("product not found")
		}
		return product, result.Error
	}
	return product, nil
}

func SaveProduct(product *entity.Product) error {
	result := config.DB.Create(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateProduct(product entity.Product) error {
	result := config.DB.Save(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteProduct(id uint) error {
	result := config.DB.Delete(&entity.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
