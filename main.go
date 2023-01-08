package main

import (
	"github.com/gin-gonic/gin"
	"pian-gin/routes"
)

func main() {
	r := gin.Default()
	routes.Routes(r)
	r.Run(":8081")
}
