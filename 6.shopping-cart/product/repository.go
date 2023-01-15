package product

import "gorm.io/gorm"

// interface
type IRepository interface {
	Save(product Product) (Product, error)
	Update(product Product) (Product, error)
	FindByProducCodeAndProductName(productCode, productName string) (Product, error)
	FindByProducCode(productCode string) (Product, error)
	DeleteByProductCode(productCode string) error
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

func (r *Repository) FindByProducCode(productCode string) (Product, error) {
	var product Product

	if err := r.db.Where("product_code = ?", productCode).Find(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r *Repository) DeleteByProductCode(productCode string) error {
	var product Product

	if err := r.db.Where("product_code = ?", productCode).Delete(&product).Error; err != nil {
		return err
	}

	return nil
}
