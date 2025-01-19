package service

import (
	"errors"
	"myapp/entity"
	"myapp/repository"
)

type ProductService interface {
	GetAllProducts() ([]entity.Product, error)
	GetProductByID(id uint) (*entity.Product, error)
	CreateProduct(product *entity.Product) error
	UpdateProduct(id uint, updatedProduct *entity.Product) error
	DeleteProduct(id uint) error
}

type productServiceImpl struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productServiceImpl{productRepository: productRepository}
}

func (s *productServiceImpl) GetAllProducts() ([]entity.Product, error) {
	return s.productRepository.GetAll()
}

func (s *productServiceImpl) GetProductByID(id uint) (*entity.Product, error) {
	product, err := s.productRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (s *productServiceImpl) CreateProduct(product *entity.Product) error {
	return s.productRepository.Save(product)
}

func (s *productServiceImpl) UpdateProduct(id uint, updatedProduct *entity.Product) error {
	product, err := s.productRepository.FindByID(id)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("product not found")
	}

	product.Name = updatedProduct.Name
	product.Price = updatedProduct.Price
	product.Stock = updatedProduct.Stock

	return s.productRepository.Update(product)
}

func (s *productServiceImpl) DeleteProduct(id uint) error {
	product, err := s.productRepository.FindByID(id)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("product not found")
	}
	return s.productRepository.Delete(id)
}
