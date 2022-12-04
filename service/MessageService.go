package service

import (
	"context"
	"fmt"
	"ginwayadi/config"
	"ginwayadi/dto"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

// Message Service

type MessageService struct{}

func (s MessageService) SendMessage(req *dto.ReqSend) {

	//send
	recipient, ok := config.ParseJID(req.No)
	if !ok {
		return
	}
	msg := &waProto.Message{Conversation: proto.String(req.Text)}
	resp, err := config.GetClient().SendMessage(context.Background(), recipient, "", msg)
	if err != nil {
		fmt.Println("Error sending message: %v", err)
	} else {
		fmt.Println("Message sent (server timestamp: %s)", resp.Timestamp)
	}
}
