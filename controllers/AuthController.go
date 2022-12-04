package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//AuthController ...
type AuthController struct{}

func (ctl AuthController) Valid(c *gin.Context) {

	// Get the Basic Authentication credentials
	username, password, _ := c.Request.BasicAuth()
	if username != os.Getenv("USERNAME_BA") || password != os.Getenv("PASSWORD_BA") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Cannot Access"})
		return
	}
}
