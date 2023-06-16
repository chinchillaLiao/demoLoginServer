package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	UserID     string `gorm:"size:22;primaryKey;"`
	Plaintext  string `gorm:"-:all"`
	CipherText string `gorm:"size:60"`
	UpdatedAt  time.Time
}

func (p *Password) Encrypt() error {
	fmt.Println(time.Now())
	// bcrypt 超慢
	CipherBytes, err := bcrypt.GenerateFromPassword([]byte(p.Plaintext), 0)
	fmt.Println(time.Now())
	p.CipherText = string(CipherBytes)
	fmt.Println(time.Now())
	return err
}

func (p *Password) Match() bool {
	fmt.Println(p.CipherText)
	fmt.Println(p.Plaintext)
	if p.Plaintext == "" {
		return false
	}
	// bcrypt 超慢
	// p.Plaintext 如果是 nil 或空字串，居然可以通過，這應該算bug吧 ?
	err := bcrypt.CompareHashAndPassword([]byte(p.CipherText), []byte(p.Plaintext))
	if err != nil {
		fmt.Println(err)
		fmt.Println(p)
		return false
	}
	return true
}
