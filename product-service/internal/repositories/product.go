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

func (r *ProductRepository) GetProducts(name, sortBy, order string) ([]models.Product, error) {
	var products []models.Product
	query := r.db.Preload("Images")
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")

	}
	if sortBy != "" {
		if order != "asc" && order != "desc" {
			order = "asc"
		}
		query = query.Order(sortBy + " " + order)
	}
	err := query.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) DeleteImage(productId, imageId uint) error {
	return r.db.Where("id = ? AND product_id = ?", imageId, productId).Delete(&models.Image{}).Error
}

func (r *ProductRepository) GetImageById(productId, imageId uint) (*models.Image, error) {
	var images models.Image
	err := r.db.Where("id = ? AND product_id = ?", imageId, productId).First(&images).Error
	if err != nil {
		return nil, err
	}
	return &images, nil
}

func (r *ProductRepository) CreateImage(image *models.Image) error {
	return r.db.Create(image).Error
}
