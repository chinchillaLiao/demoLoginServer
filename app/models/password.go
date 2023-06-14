package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	UserID     string `gorm:"size:22;primaryKey"`
	Plaintext  string `gorm:"-:all"`
	CipherText string `gorm:"size:60"`
	UpdatedAt  time.Time
}

func (p *Password) Encrypt() error {
	CipherBytes, err := bcrypt.GenerateFromPassword([]byte(p.Plaintext), 14)
	p.CipherText = string(CipherBytes)
	return err
}

func (p *Password) Match() bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.CipherText), []byte(p.Plaintext))
	if err != nil {
		fmt.Println(err)
		fmt.Println(p)
		return false
	}
	return true
}
