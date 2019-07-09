package models

import (
	"log"
	"time"
)

type Message struct {
	ID        int        `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	UserID    int        `json:"user_id"`
	Content   string     `json:"content"`
}

func (message *Message) Validate() bool {
	if message.Content == "" {
		log.Fatalln("Message content is empty.")
		return false
	}

	return true
}

func (message *Message) Create() uint {
	if ok := message.Validate(); !ok {
		return 0
	}

	GetDB().Create(message)

	if message.ID <= 0 {
		log.Fatalln("Error creating message")
		return 0
	}

	return message.ID
}

func GetMessagesByUsername(username string) []*Message {
	messages := make([]*Message, 0)
	user := User{}

	err := GetDB().Table("users").Where("user_name = ?", username).First(&user).Error
	if err != nil {
		log.Fatalf("Error retrieving messages for user: %v\n", err)
		return nil
	}

	err = GetDB().Table("messages").Where("user_id = ?", user.ID).Find(&messages).Error

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
