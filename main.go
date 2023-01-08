package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"pian-gin/routes"
)

func main() {
	r := gin.Default()
	routes.BurgerRoute(r)
	log.Fatal(r.Run(":8081"))
}
