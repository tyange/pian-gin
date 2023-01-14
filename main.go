package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"pian-gin/config"
	"pian-gin/routes"
)

func main() {
	r := gin.Default()
	config.ConnectDB()
	routes.BurgerRoute(r)
	routes.UserRouter(r)
	log.Fatal(r.Run(":8081"))
}
