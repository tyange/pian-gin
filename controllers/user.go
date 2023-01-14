package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"pian-gin/config"
	"pian-gin/models"
	"time"
)

func SignUp(c *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(400, err)

		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}
	result := config.DB.Create(&user)

	if result.Error != nil {
		c.JSON(400, result.Error)

		return
	}

	c.JSON(200, &user)
}

func SignIn(c *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	user := models.User{}
	config.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(400, gin.H{
			"error": "Invalid email or password",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Pasword",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(400, err)

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(200, gin.H{})
}
