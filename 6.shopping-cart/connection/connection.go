package connection

import (
	"errors"
	"fmt"

	"telkom-tect-test/6.shopping-cart/product"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection(env map[string]string) (*gorm.DB, error) {
	// ammbil value env
	dbUsername := env["dbUsername"]
	dbPassword := env["dbPassword"]
	dbHost := env["dbHost"]
	dbPort := env["dbPort"]
	dbName := env["dbName"]

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
