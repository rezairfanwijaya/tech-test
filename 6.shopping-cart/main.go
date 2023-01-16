package main

import (
	"log"
	"telkom-tect-test/6.shopping-cart/connection"
	"telkom-tect-test/6.shopping-cart/route"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := connection.Connection(".env")
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
