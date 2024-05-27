package repositories

import (
	"gorm.io/gorm"
	"product-service/internal/models"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Images").First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) UpdateProduct(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) DeleteProduct(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

func (r *ProductRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Images").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
