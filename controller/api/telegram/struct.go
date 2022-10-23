package telegram

import "github.com/semmisha/ClientAPI"

type TelegramStruct struct {
	Message string
	Config  map[string]string
	Client  *ClientAPI.Message
}

func NewTelegramStruct() *TelegramStruct {
	return &TelegramStruct{}
}
