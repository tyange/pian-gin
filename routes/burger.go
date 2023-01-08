package routes

import "github.com/gin-gonic/gin"

func BurgerRoute(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(200, "this burger route.")
	})
}
