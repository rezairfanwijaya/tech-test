package product

// interface
type IService interface {
	Create(input ProductInput) (Product, error)
}

// struct untuk dependency dan set method
type Service struct {
	repoProduct IRepository
}

// new service
func NewService(repoProduct IRepository) *Service {
	return &Service{repoProduct}
}

func (s *Service) Create(input ProductInput) (Product, error) {
	// mapping
	var product Product
	product.ProductCode = input.ProductCode
	product.ProductName = input.ProductName
	product.Quantity = input.Quantity

	// jika product sudah ada didalam cart maka
	// tambahakan kuantitasnya
	productExist, err := s.repoProduct.FindByProducCodeAndProductName(input.ProductCode, input.ProductName)
	if err != nil {
		return productExist, err
	}

	// jika ada tambahkan kuantitasnya
	if productExist.ID != 0 {
		productExist.Quantity += input.Quantity
		productUpdate, err := s.repoProduct.Update(productExist)
		if err != nil {
			return productUpdate, err
		}

		return productUpdate, nil
	} else {
		newProduct, err := s.repoProduct.Save(product)
		if err != nil {
			return newProduct, err
		}

		return newProduct, nil
	}
}
