package repository

import (
	"errors"
	"myapp/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll() ([]entity.Product, error)
	FindByID(id uint) (*entity.Product, error)
	Save(product *entity.Product) error
	Update(product *entity.Product) error
	Delete(id uint) error
}

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryImpl{db: db}
}

func (r *productRepositoryImpl) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepositoryImpl) FindByID(id uint) (*entity.Product, error) {
	var product entity.Product
	if err := r.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepositoryImpl) Save(product *entity.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepositoryImpl) Update(product *entity.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entity.Product{}, id).Error
}
