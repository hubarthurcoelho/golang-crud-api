package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hubarthurcoelho/golang-crud-api/config"
	"github.com/hubarthurcoelho/golang-crud-api/database"
	"github.com/hubarthurcoelho/golang-crud-api/router"
)

func SetUpAndRunApp() error {
	envErr := config.LoadEnvVariables()
	if envErr != nil {
		return envErr
	}

	dbErr := database.ConnectToDatabase()
	if dbErr != nil {
		return dbErr
	}

	r := gin.Default()

	router.SetupRoutes(r)

	r.Run()

	return nil
}
