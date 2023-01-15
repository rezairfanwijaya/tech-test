package product

import (
	"errors"
	"fmt"
)

// interface
type IService interface {
	Create(input ProductInput) (Product, error)
	Delete(input DeleteProductInput) error
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

func (s *Service) Delete(input DeleteProductInput) error {
	// cari apakah ada product yang akan di delete
	productByProductCode, err := s.repoProduct.FindByProducCode(input.ProductCode)
	if err != nil {
		return err
	}

	// jika id lebih dari nol berarti ada
	// dan lakukan delete
	if productByProductCode.ID != 0 {
		if err := s.repoProduct.DeleteByProductCode(productByProductCode.ProductCode); err != nil {
			return err
		}
		return nil
	}

	// jika tidak ada
	errMsg := fmt.Sprintf("product code %v not found", input.ProductCode)
	return errors.New(errMsg)
}
