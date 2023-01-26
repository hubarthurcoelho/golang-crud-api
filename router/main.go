package router

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	userRouter(r)
}
