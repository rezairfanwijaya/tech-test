package product

// model product
type Product struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	ProductCode string `json:"product_code"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}
