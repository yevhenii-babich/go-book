package models

import "gorm.io/gorm"

// User - модель користувача
type User struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

type Profile struct {
	gorm.Model
	UserID   uint   `json:"user_id" gorm:"unique;not null"`
	FullName string `json:"full_name" gorm:"not null"`
	Age      int    `json:"age" gorm:"default:30"`
}

// ProfileWithUserEmail - модель профілю з електронною адресою користувача
type ProfileWithUserEmail struct {
	Profile
	Email string `json:"email"`
}

// ProfileWithUser - модель профілю з користувачем
type ProfileWithUser struct {
	Profile
	User User `json:"user" gorm:"foreignKey:UserID;references:ID"`
}
