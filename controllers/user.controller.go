package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hubarthurcoelho/golang-crud-api/commands"
	"github.com/hubarthurcoelho/golang-crud-api/models"
	"github.com/hubarthurcoelho/golang-crud-api/validations"
)

func CreateUser(c *gin.Context) {
	var body models.CreateUserData
	c.BindJSON(&body)

	if errs := validations.ValidateBody(body); errs != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errs.Error())
		return
	}

	user := models.User{Username: body.Username, Email: body.Email, Password: body.Password}

	if err := commands.CreateUserCommand(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}
