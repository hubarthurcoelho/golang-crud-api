package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hubarthurcoelho/golang-crud-api/config"
	"github.com/hubarthurcoelho/golang-crud-api/database"
	"github.com/hubarthurcoelho/golang-crud-api/database/migrations"
	"github.com/hubarthurcoelho/golang-crud-api/router"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpAndRunApp() error {
	envErr := config.LoadEnvVariables()
	if envErr != nil {
		return envErr
	}

	r := gin.Default()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.SetupRoutes(r)

	r.Run()

	DB, dbErr := database.ConnectToDatabase()
	if dbErr != nil {
		return dbErr
	}

	migrateErr := migrations.MigrateDB(DB)
	if migrateErr != nil {
		return migrateErr
	}

	return nil
}
