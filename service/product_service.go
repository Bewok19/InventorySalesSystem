package service

import (
	"myapp/entity"
	"myapp/repository"
)

func GetAllProducts() ([]entity.Product, error) {
	products, err := repository.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func CreateProduct(product entity.Product) error {
	err := repository.SaveProduct(&product)
	if err != nil {
		return err
	}
	return nil
}

func GetProductByID(id int) (*entity.Product, error) {
	product, err := repository.GetProductByID(uint(id))
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func UpdateProduct(id int, updatedProduct entity.Product) error {
	// Cek apakah produk ada
	product, err := repository.GetProductByID(uint(id))
	if err != nil {
		return err
	}

	// Update data produk
	product.Name = updatedProduct.Name
	product.Price = updatedProduct.Price
	product.Stock = updatedProduct.Stock

	err = repository.UpdateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(id int) error {
	err := repository.DeleteProduct(uint(id))
	if err != nil {
		return err
	}
	return nil
}
