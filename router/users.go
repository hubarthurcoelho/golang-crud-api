package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/hubarthurcoelho/golang-crud-api/controllers"
)

func userRouter(r *gin.Engine) {
	users := r.Group("/users")
	users.POST("/", controller.CreateUser)
}
