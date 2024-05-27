package services

import (
	"mime/multipart"
	"product-service/internal/models"
	"product-service/internal/repositories"
	"product-service/utils"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductService(productRepo *repositories.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (s *ProductService) CreateProduct(input *models.CreateProductInput, file []*multipart.FileHeader) (*models.Product, error) {
	product := &models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       *input.Stock,
	}
	for _, f := range file {
		imageURL, err := utils.UploadFile(f)
		if err != nil {
			return nil, err
		}
		product.Images = append(product.Images, models.Image{
			URL: imageURL,
		})

	}
	err := s.productRepo.CreateProduct(product)
	return product, err
}

func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	return s.productRepo.GetProductByID(id)
}

func (s *ProductService) UpdateProduct(id uint, input *models.UpdateProductInput) (*models.Product, error) {
	product, err := s.productRepo.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Stock = *input.Stock

	err = s.productRepo.UpdateProduct(product)
	return product, err
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.productRepo.DeleteProduct(id)
}
