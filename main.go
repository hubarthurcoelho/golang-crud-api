package main

import (
	"log"

	"github.com/hubarthurcoelho/golang-crud-api/app"
)

// @title Swagger xd
func main() {
	err := app.SetUpAndRunApp()
	if err != nil {
		log.Fatal(err)
	}

}
