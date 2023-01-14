package product

type ProductInput struct {
	ProductCode string `json:"product_code" binding:"required"`
	ProductName string `json:"product_name" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required,min=1"`
}
