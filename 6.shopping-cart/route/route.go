package route

import (
	"telkom-tect-test/6.shopping-cart/handler"
	"telkom-tect-test/6.shopping-cart/product"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRoute(db *gorm.DB, route *gin.Engine) {
	// repo product
	repoProduct := product.NewRepository(db)
	// service product
	serviceProduct := product.NewService(repoProduct)
	// handler product
	hanlderProduct := handler.NewHandlerProduct(serviceProduct)

	route.POST("/product", hanlderProduct.CreateNewProduct)
}
