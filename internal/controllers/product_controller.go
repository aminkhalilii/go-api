package controllers

import (
	"go-api/internal/models"
	"go-api/internal/services"
	"go-api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService services.ProductServiceInterface
}

func NewProductController(productService services.ProductServiceInterface) *ProductController {
	return &ProductController{productService: productService}
}

func (pc *ProductController) GetAllProducts(c *gin.Context) {
	products, err := pc.productService.GetAllProducts()
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, "products retrieved successfully", products)
}

func (pc *ProductController) GetProductByID(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid product ID")
		return
	}
	product, err := pc.productService.GetProductByID(id)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())

		return
	}
	if product == nil {
		utils.Error(c, http.StatusNotFound, "Product not found")
		return
	}
	utils.Success(c, http.StatusOK, "product retrieved successfully", product)
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid input")
		return
	}
	newProduct, err := pc.productService.CreateProduct(&product)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusCreated, "Product created successfully", newProduct)
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var product models.Product
	err = c.ShouldBindJSON(&product)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	newProduct, err := pc.productService.UpdateProduct(id, &product)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, "Product updated successfully", newProduct)
}
func (pc *ProductController) DeleteProduct(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid Product ID")
		return
	}
	err = pc.productService.DeleteProduct(id)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, "Product deleted successfully", nil)

}
