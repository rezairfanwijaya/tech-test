package product

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// interface
type IService interface {
	Create(input ProductInput) (Product, error)
	Delete(input DeleteProductInput) error
	GetAll(params map[string]string) ([]Product, int, error)
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
	productByProductCode, err := s.repoProduct.FindByProductCode(input.ProductCode)
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

func (s *Service) GetAll(params map[string]string) ([]Product, int, error) {
	// cek apakah ada params nama produk atau kuantitas
	productName := params["product_name"]
	productQuantity := params["product_quantity"]

	// jika ada kedua parameter
	if productName != "" && productQuantity != "" {
		// cek apakah ada product name
		productByProductName, httpCode, err := s.getProductByProductName(productName)
		if err != nil {
			return productByProductName, httpCode, err
		}

		// cek apakah ada product quantity
		productByProductQuantity, httpCode, err := s.getProductByProductQuantity(productQuantity)
		if err != nil {
			return productByProductQuantity, httpCode, err
		}

		quantity, _ := s.convertStringToInteger(productQuantity)

		// jika ada semua maka lakukan query keduanya
		products, err := s.repoProduct.FindByProducNameAndProductQuantity(productName, quantity)
		if err != nil {
			return products, http.StatusNotFound, err
		}

		log.Println(products)

		if len(products) == 0 {
			errMsg := fmt.Sprintf("product with quantity %v and name %v not found", quantity, productName)
			return products, http.StatusNotFound, errors.New(errMsg)
		}

		return products, http.StatusOK, nil
	}


	// jika hanya ada parameter product name
	if productName != "" {
		productByProductName, httpCode, err := s.getProductByProductName(productName)
		if err != nil {
			return productByProductName, httpCode, err
		}

		return productByProductName, httpCode, nil
	}

	// jika hanya ada parameter quantity
	if productQuantity != "" {
		productByProductQuantity, httpCode, err := s.getProductByProductQuantity(productQuantity)
		if err != nil {
			return productByProductQuantity, httpCode, err
		}

		return productByProductQuantity, httpCode, nil
	}

	// jika tidak ada params
	allProducts, err := s.repoProduct.FindAll()
	if err != nil {
		return allProducts, http.StatusNotFound, err
	}

	return allProducts, http.StatusOK, nil
}

func (s *Service) getProductByProductName(productName string) ([]Product, int, error) {
	productByProductName, err := s.repoProduct.FindByProductName(productName)
	if err != nil {
		return productByProductName, http.StatusNotFound, err
	}

	// cek apakah ada product atau tidak
	if len(productByProductName) == 0 {
		errMsg := fmt.Sprintf("product name %v not found", productName)
		return productByProductName, http.StatusNotFound, errors.New(errMsg)
	}

	return productByProductName, http.StatusOK, nil
}

func (s *Service) getProductByProductQuantity(productQuantity string) ([]Product, int, error) {
	// convert quantity ke integer
	quantity, err := s.convertStringToInteger(productQuantity)
	if err != nil {
		errMsg := fmt.Sprint("quantity must be integer and grather than 0")
		return []Product{}, http.StatusBadRequest, errors.New(errMsg)
	}

	productByProductQuantity, err := s.repoProduct.FindByProductQuantity(quantity)
	if err != nil {
		return productByProductQuantity, http.StatusNotFound, err
	}

	// cek apakah ada product atau tidak
	if len(productByProductQuantity) == 0 {
		errMsg := fmt.Sprintf("product quantity %v not found", productQuantity)
		return productByProductQuantity, http.StatusNotFound, errors.New(errMsg)
	}

	return productByProductQuantity, http.StatusOK, nil
}

func (s *Service) convertStringToInteger(value string) (int, error) {
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return valueInt, err
	}

	return valueInt, nil
}
