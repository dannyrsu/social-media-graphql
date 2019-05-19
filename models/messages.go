package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model
	UserID  int
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
	if !ok := message.Validate() {
		return false
	}

	GetDB().Create(message)

	if message.ID <= 0 {
		log.Fatalln("Error creating message")
		return false
	}

	return true
}

func GetMessage(messageID uint) *Message {
	message := &Message{}

	err := GetDB().Table("messages").Where("id = ?", id).Find(message).Error

	if err != nil {
		log.Fatalf("Error retreiving message: %v\n", err)
		return nil
	}

	return message
}

func GetMessages(userID uint) []*Message {
	messages := make([]*Message, 0)

	err := GetDB().Table("messages").Where("user_id = ?", UserID).Find(&messages).Error

	if err != nil {
		log.Fatalf("Error retreiving messages for user: %v", err)
		return nil
	}

	return messages
}