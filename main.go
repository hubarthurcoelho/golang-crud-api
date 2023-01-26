package main

import (
	"log"

	"github.com/hubarthurcoelho/golang-crud-api/app"
)

func main() {
	err := app.SetUpAndRunApp()
	if err != nil {
		log.Fatal(err)
	}
}
