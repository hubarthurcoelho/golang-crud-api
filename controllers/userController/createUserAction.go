package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hubarthurcoelho/golang-crud-api/commands"
	"github.com/hubarthurcoelho/golang-crud-api/models"
	"github.com/hubarthurcoelho/golang-crud-api/validations"
)

func CreateUserAction(c *gin.Context) {
	var body models.User
	c.BindJSON(&body)

	if errs := validations.ValidateBody(body); errs != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errs.Error())
		return
	}

	if err := commands.CreateUserCommand(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": body})
}
