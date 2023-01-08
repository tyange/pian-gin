package controllers

import "github.com/gin-gonic/gin"

func BurgerController(c *gin.Context) {
	c.String(200, "this is burger route")
}
