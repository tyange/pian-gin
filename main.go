package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"pian-gin/routes"
)

func main() {
	r := gin.Default()
	routes.Routes(r)
	log.Fatal(r.Run(":8081"))
}
