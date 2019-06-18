package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
}

func (message *Message) Validate() bool {
	if message.Content == "" {
		log.Fatalln("Message content is empty.")
		return false
	}

	return true
}

func (message *Message) Create() bool {
	if ok := message.Validate(); !ok {
		return false
	}

	GetDB().Create(message)

	if message.ID <= 0 {
		log.Fatalln("Error creating message")
		return false
	}

	return true
}

func GetMessagesByUsername(username string) []*Message {
	messages := make([]*Message, 0)

	err := GetDB().Table("messages").Where("user_name = ?", username).Find(&messages).Error

	if err != nil {
		log.Fatalf("Error retrieving messages for user: %v\n", err)
		return nil
	}

	return messages
}

func GetAllMessages() []*Message {
	messages := make([]*Message, 0)

	err := GetDB().Table("messages").Find(&messages).Error

	if err != nil {
		log.Fatalf("Error retrieving messages: %v\n", err)
		return nil
	}

	return messages
}
