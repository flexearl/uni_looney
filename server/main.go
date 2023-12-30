package main

import (
	"log"
	"net/http"

	"github.com/flexearl/uni_looney.git/routes"
)

func main() {

	routes.HandleRoutes()
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
