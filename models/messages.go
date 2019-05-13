package models

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model
	UserID  int
	Content string `json:"content"`
}
