package helper

import (
	"errors"
	"fmt"

	"telkom-tect-test/6.shopping-cart/product"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type responseAPI struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type productResponse struct {
	ProductCode string `json:"product_code"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}

type ResponseNil struct{}

func FormatProduct(product product.Product) productResponse {
	return productResponse{
		ProductCode: product.ProductCode,
		ProductName: product.ProductName,
		Quantity:    product.Quantity,
	}
}

func FormatProducts(products []product.Product) []productResponse {
	var result []productResponse

	for _, product := range products {
		singleProductFormated := FormatProduct(product)
		result = append(result, singleProductFormated)
	}

	return result
}

func SetResponseAPI(status, message string, code int, data interface{}) responseAPI {
	meta := meta{
		Status:  status,
		Code:    code,
		Message: message,
	}

	return responseAPI{
		Meta: meta,
		Data: data,
	}
}

func SetErrorBinding(err error) []string {
	var errBinding []string

	for _, v := range err.(validator.ValidationErrors) {
		errMsg := fmt.Sprintf("error on field : %v, condition : %v", v.Field(), v.ActualTag())
		errBinding = append(errBinding, errMsg)
	}

	return errBinding
}

func GetEnv(path string) (map[string]string, error) {
	env, err := godotenv.Read(path)

	if err != nil {
		errMsg := fmt.Sprintf("ERR GET ENV : %v", err)
		return env, errors.New(errMsg)
	}

	return env, nil
}
