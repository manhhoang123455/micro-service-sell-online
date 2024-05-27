package controllers

import (
	"net/http"
	"product-service/internal/models"
	"product-service/internal/services"
	"product-service/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *services.ProductService
}

func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with details
// @Tags product
// @Accept json
// @Produce json
// @Param product body models.CreateProductInput true "Product"
// @Success 200 {object} models.Product
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /products [post]
func (pc *ProductController) CreateProduct(c *gin.Context) {
	var input models.CreateProductInput
	if err := c.ShouldBind(&input); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	form, _ := c.MultipartForm()
	files := form.File["images"]

	product, err := pc.productService.CreateProduct(&input, files)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create product")
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetProductByID godoc
// @Summary Get a product by ID
// @Description Get product details by ID
// @Tags product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /products/{id} [get]
func (pc *ProductController) GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := pc.productService.GetProductByID(uint(id))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Product not found")
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update product details
// @Tags product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.UpdateProductInput true "Product"
// @Success 200 {object} models.Product
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /products/{id} [put]
func (pc *ProductController) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var input models.UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	product, err := pc.productService.UpdateProduct(uint(id), &input)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to update product")
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product by ID
// @Tags product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /products/{id} [delete]
func (pc *ProductController) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	err = pc.productService.DeleteProduct(uint(id))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Product not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
