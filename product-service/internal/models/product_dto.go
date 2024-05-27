package models

type CreateProductInput struct {
	Name        string  `form:"name" binding:"required"`
	Description string  `form:"description"`
	Price       float64 `form:"price" binding:"required"`
	Stock       *int    `form:"stock" binding:"required"`
}

type UpdateProductInput struct {
	Name        string  `form:"name" binding:"required"`
	Description string  `form:"description" binding:"required"`
	Price       float64 `form:"price" binding:"required"`
	Stock       *int    `form:"stock" binding:"required"`
}
