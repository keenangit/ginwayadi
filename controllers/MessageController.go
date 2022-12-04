package controllers

import (
	"fmt"
	"ginwayadi/dto"
	"ginwayadi/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageController struct{}

var messageService = new(service.MessageService)
var reqSend = new(dto.ReqSend)

func (ctrl MessageController) Send(c *gin.Context) {
	fmt.Printf("ClientIP: %s\n", c.ClientIP())

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&reqSend); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Invalid Request"})
	}

	messageService.SendMessage(reqSend)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully"})

}
