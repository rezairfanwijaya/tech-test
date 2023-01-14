package product

import "gorm.io/gorm"

// interface
type IRepository interface {
	Save(product Product) (Product, error)
	Update(product Product) (Product, error)
	FindByProducCodeAndProductName(productCode, productName string) (Product, error)
}

// struct untuk dependency dan set method
type Repository struct {
	db *gorm.DB
}

// new repo
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

// implementasi method
func (r *Repository) Save(product Product) (Product, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r *Repository) Update(product Product) (Product, error) {
	if err := r.db.Save(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r *Repository) FindByProducCodeAndProductName(productCode, productName string) (Product, error) {
	var product Product

	if err := r.db.Where("product_code = ? and product_name = ?", productCode, productName).Find(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}
