package models

import (
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type Token struct {
	UserID uint
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	UserName  string `json:"username"`
	Password  string `json:password`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Token     string `json:"token"; sql:"-"`
}

func (user *User) Validate() bool {
	if !strings.Contains(user.Email, "@") {
		return false
	}

	if len(user.Password) < 8 {
		return false
	}
	tempUser := &User{}
	err := GetDB().Table("users").Where("email = ?", tempUser.Email).First(tempUser).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		log.Fatalf("Connection Error: %v\n", err)
		return false
	}

	if tempUser.Email != "" {
		log.Fatalln("Email address already in use by another user.")
		return false
	}

	return true
}

func (user *User) Create() uint {
	if !user.Validate() {
		log.Fatalln("Can not create user account.")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	if user.ID <= 0 {
		log.Fatalln("Error creating new user: %v\n")
		return 0
	}

	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString
	user.Password = ""

	return user.ID
}

func Login(email, password string) bool {
	user := &User{}

	err := GetDB().Table("users").Where("email = ?", email).First(user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Fatalln("Email address not found")
		} else {
			log.Fatalf("Connection error: %v\n", err)
		}

		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Fatalf("Invalid credentials : %v\n", err)
		return false
	}

	user.Password = ""
	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	return true
}
