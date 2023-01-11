package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pian-gin/config"
	"pian-gin/models"
)

func GetAllBurger(c *gin.Context) {
	var burgers []models.Burger
	result := config.DB.Find(&burgers)

	if result.Error != nil {
		c.JSON(400, result.Error)
	}

	c.JSON(200, &burgers)
}

func GetBurger(c *gin.Context) {
	fmt.Println(c.Param("id"))
}

func AddBurger(c *gin.Context) {
	burger := models.Burger{}
	err := c.Bind(&burger)

	if err != nil {
		c.JSON(400, err)
	}

	result := config.DB.Create(&burger)

	if result.Error != nil {
		c.JSON(400, result.Error)
	}

	c.JSON(200, gin.H{
		"burgerId": burger.Id,
	})
}
