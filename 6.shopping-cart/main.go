package main

import (
	"log"
	"os"
	"telkom-tect-test/6.shopping-cart/connection"
)

func main() {
	// mengambil env
	dbUsername := os.Getenv("USERNAME")
	dbPassword := os.Getenv("PASSWORD")
	dbHost := os.Getenv("HOST")
	dbPort := os.Getenv("PORT")
	dbName := os.Getenv("DBNAME")

	mapEnv := map[string]string{
		"dbUsername": dbUsername,
		"dbPassword": dbPassword,
		"dbHost":     dbHost,
		"dbPort":     dbPort,
		"dbName":     dbName,
	}

	_, err := connection.Connection(mapEnv)
	if err != nil {
		log.Fatal(err)
		return
	}

}
