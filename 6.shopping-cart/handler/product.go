package handler

import (
	"net/http"
	"telkom-tect-test/6.shopping-cart/helper"
	"telkom-tect-test/6.shopping-cart/product"

	"github.com/gin-gonic/gin"
)

// struct dependency
type productHandler struct {
	serviceProduct product.IService
}

// new handler
func NewHandlerProduct(serviceProduct product.IService) *productHandler {
	return &productHandler{serviceProduct}
}

// implementasi
func (h *productHandler) CreateNewProduct(c *gin.Context) {
	var input product.ProductInput

	// binding
	if err := c.ShouldBindJSON(&input); err != nil {
		errMsg := helper.SetErrorBinding(err)
		response := helper.SetResponseAPI(
			"failed",
			"failed to save new product",
			http.StatusBadRequest,
			errMsg,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// panggil service
	newProduct, err := h.serviceProduct.Create(input)
	if err != nil {
		response := helper.SetResponseAPI(
			"failed",
			"failed to save new product",
			http.StatusBadRequest,
			err.Error(),
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.SetResponseAPI(
		"success",
		"success to save new product",
		http.StatusCreated,
		helper.FormatProduct(newProduct),
	)

	c.JSON(http.StatusCreated, response)
}

func (h *productHandler) DeleteByProductCode(c *gin.Context) {
	var input product.DeleteProductInput

	// binding
	if err := c.ShouldBindJSON(&input); err != nil {
		errMsg := helper.SetErrorBinding(err)
		response := helper.SetResponseAPI(
			"failed",
			"failed to delete product",
			http.StatusBadRequest,
			errMsg,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// panggil service
	if err := h.serviceProduct.Delete(input); err != nil {
		response := helper.SetResponseAPI(
			"failed",
			"failed to delete product",
			http.StatusNotFound,
			err.Error(),
		)

		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.SetResponseAPI(
		"success",
		"success delete product",
		http.StatusOK,
		helper.ResponseNil{},
	)

	c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetAllProduct(c *gin.Context) {
	// ambil query params
	productName := c.Query("name")
	productQuantity := c.Query("quantity")

	mapParams := map[string]string{
		"product_name":     productName,
		"product_quantity": productQuantity,
	}

	// panggil service
	products, httpCode, err := h.serviceProduct.GetAll(mapParams)
	if err != nil {
		response := helper.SetResponseAPI(
			"failed",
			"failed to get product",
			httpCode,
			err.Error(),
		)

		c.JSON(httpCode, response)
		return
	}

	response := helper.SetResponseAPI(
		"success",
		"success get product",
		httpCode,
		helper.FormatProducts(products),
	)

	c.JSON(httpCode, response)
}
