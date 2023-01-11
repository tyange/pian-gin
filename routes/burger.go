package routes

import (
	"github.com/gin-gonic/gin"
	"pian-gin/controllers"
)

func BurgerRoute(router *gin.Engine) {
	router.GET("/burger", controllers.GetAllBurger)
	router.POST("/burger", controllers.AddBurger)
	router.GET("/burger/:id", controllers.GetBurger)
}
