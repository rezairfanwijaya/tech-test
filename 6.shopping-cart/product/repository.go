package product

import "gorm.io/gorm"

// interface
type IRepository interface {
	Save(product Product) (Product, error)
	Update(product Product) (Product, error)
	FindByProducCodeAndProductName(productCode, productName string) (Product, error)
	FindByProducNameAndProductQuantity(productName string, productQuantity int) ([]Product, error)
	FindByProductCode(productCode string) (Product, error)
	FindByProductName(productName string) ([]Product, error)
	FindByProductQuantity(productQuantity int) ([]Product, error)
	DeleteByProductCode(productCode string) error
	FindAll() ([]Product, error)
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

func (r *Repository) FindByProductName(productName string) ([]Product, error) {
	var product []Product

	if err := r.db.Where("product_name = ?", productName).Find(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r *Repository) FindByProductQuantity(productQuantity int) ([]Product, error) {
	var product []Product

	if err := r.db.Where("quantity = ?", productQuantity).Find(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r *Repository) FindByProducNameAndProductQuantity(productName string, productQuantity int) ([]Product, error) {
	var product []Product

	if err := r.db.Where("product_name = ? and quantity = ?", productName, productQuantity).Find(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r *Repository) FindByProductCode(productCode string) (Product, error) {
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

func (r *Repository) FindAll() ([]Product, error) {
	var product []Product

	if err := r.db.Find(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}
