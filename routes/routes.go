package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(r *gin.Engine) {
	r.GET("/burgers", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"message": "this is burgers routes",
		})
	})
}
