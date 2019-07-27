package models

import (
	"errors"
	"log"
)

type Message struct {
	BaseModel
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

func (message *Message) Create() (*Message, error) {
	if ok := message.Validate(); !ok {
		return nil, errors.New("error validating message")
	}

	GetDB().Create(message)

	if message.ID <= 0 {
		log.Fatalln("Error creating message")
		return nil, errors.New("error creating message")
	}

	return message, nil
}

func GetMessagesByEmail(email string) ([]*Message, error) {
	messages := make([]*Message, 0)
	user := User{}

	err := GetDB().Table("users").Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Fatalf("Error retrieving messages for user: %v\n", err)
		return nil, err
	}

	err = GetDB().Table("messages").Where("user_id = ?", user.ID).Find(&messages).Error

	if err != nil {
		log.Fatalf("Error retrieving messages for user: %v\n", err)
		return nil, err
	}

	return messages, nil
}

func GetAllMessages() ([]*Message, error) {
	messages := make([]*Message, 0)

	err := GetDB().Table("messages").Find(&messages).Error

	if err != nil {
		log.Fatalf("Error retrieving messages: %v\n", err)
		return nil, err
	}

	return messages, nil
}
