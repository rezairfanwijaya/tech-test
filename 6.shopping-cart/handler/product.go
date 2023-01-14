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
