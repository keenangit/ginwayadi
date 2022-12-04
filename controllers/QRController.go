package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QRController struct{}

func (ctrl QRController) GetQR(c *gin.Context) {
	fmt.Printf("ClientIP: %s\n", c.ClientIP())

	c.JSON(http.StatusOK, gin.H{"message": "Successfully"})

}
