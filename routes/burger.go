package routes

import (
	"github.com/gin-gonic/gin"
	"pian-gin/controllers"
)

func BurgerRoute(router *gin.Engine) {
	router.GET("/", controllers.BurgerController)
}
