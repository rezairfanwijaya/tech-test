package main

import (
	"log"
	"os"
	"telkom-tect-test/6.shopping-cart/connection"
	"telkom-tect-test/6.shopping-cart/route"

	"github.com/gin-gonic/gin"
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

	db, err := connection.Connection(mapEnv)
	if err != nil {
		log.Fatal(err)
		return
	}

	// intial gin
	r := gin.Default()

	// new route
	route.NewRoute(db, r)

	if err := r.Run(":8686"); err != nil {
		log.Fatal(err.Error())
	}

}
