package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	name     string
	email    string
	Password Password
}

type Password struct {
	UserID     uint
	Plaintext  string `gorm:"-:all"`
	CipherText string
}
