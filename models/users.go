package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
