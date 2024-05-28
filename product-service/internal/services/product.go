package services

import (
	"errors"
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

func (s *ProductService) GetProducts(name, sortBy, order string) ([]models.Product, error) {
	return s.productRepo.GetProducts(name, sortBy, order)
}

func (ps *ProductService) AddImage(productId uint, file *multipart.FileHeader) (*models.Image, error) {
	product, err := ps.productRepo.GetProductByID(productId)
	if err != nil {
		return nil, errors.New("product not found")
	}

	url, err := utils.UploadFile(file)
	if err != nil {
		return nil, errors.New("failed to upload image")
	}

	image := &models.Image{
		URL:       url,
		ProductID: product.ID,
	}

	if err := ps.productRepo.CreateImage(image); err != nil {
		return nil, errors.New("failed to add image to product")
	}

	return image, nil
}

func (s *ProductService) DeleteImage(productId, imageId uint) error {
	return s.productRepo.DeleteImage(productId, imageId)
}

func (s *ProductService) GetImageByProducts(productId, imageId uint) (*models.Image, error) {
	return s.productRepo.GetImageById(productId, imageId)
}
