package controllers

import (
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
	burger := models.Burger{}
	burgerId := c.Param("id")

	result := config.DB.Find(&burger, burgerId)

	if result.Error != nil {
		c.JSON(400, result.Error)
	}

	c.JSON(200, &burger)
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
		"burgerId": burger.ID,
	})
}

func EditBurger(c *gin.Context) {
	burger := models.Burger{}
	foundBurger := config.DB.Find(&burger, c.Param("id"))

	if foundBurger.Error != nil {
		c.JSON(400, foundBurger.Error)
	}

	err := c.Bind(&burger)

	if err != nil {
		c.JSON(400, err)
	}

	result := config.DB.Save(&burger)

	if result.Error != nil {
		c.JSON(400, result.Error)
	}

	c.JSON(200, &burger)
}
