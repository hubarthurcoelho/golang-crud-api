package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hubarthurcoelho/golang-crud-api/controllers/userController"
)

func userRouter(r *gin.Engine) {
	users := r.Group("/users")
	users.POST("/", userController.CreateUserAction)
}
