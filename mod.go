package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	name  string
	email string
}

type Password struct {
	User       User   `gorm:"primaryKey"`
	Plaintext  string `gorm:"-:all"`
	CipherText string
}
