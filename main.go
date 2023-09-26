package main

import (
	"fmt"
	"log"

	"github.com/hubarthurcoelho/go-errs/errs"
	"github.com/hubarthurcoelho/golang-crud-api/app"
)

// @title Swagger xd
func main() {
	err := app.SetUpAndRunApp()
	if err != nil {
		log.Fatal(errs.E(err))
	}

	fmt.Println("started")
}
