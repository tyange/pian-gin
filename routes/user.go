package routes

import (
	"github.com/gin-gonic/gin"
	"pian-gin/controllers"
)

func UserRouter(router *gin.Engine) {
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.SignIn)
}
