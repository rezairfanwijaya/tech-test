package connection

import (
	"errors"
	"fmt"

	"telkom-tect-test/6.shopping-cart/helper"
	"telkom-tect-test/6.shopping-cart/product"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection(path string) (*gorm.DB, error) {
	// ammbil env
	env, err := helper.GetEnv(path)
	if err != nil {
		return nil, err
	}

	dbUsername := env["DATABASE_USERNAME"]
	dbPassword := env["DATABASE_PASSWORD"]
	dbHost := env["DATABASE_HOST"]
	dbPort := env["DATABASE_PORT"]
	dbName := env["DATABASE_NAME"]

	// siapkan dns
	dsn := fmt.Sprintf("%s:%s@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	// lakukan koneksi
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		errMsg := fmt.Sprintf("ERR CONNECTION : %v", err.Error())
		return db, errors.New(errMsg)
	}

	// migration schema
	if err := db.AutoMigrate(&product.Product{}); err != nil {
		errMsg := fmt.Sprintf("ERR MIGRATION : %v", err.Error())
		return db, errors.New(errMsg)
	}

	return db, nil
}
