package main

import "golang.org/x/crypto/bcrypt"

func (p *Password) Encrypt() error {
	CipherBytes, err := bcrypt.GenerateFromPassword([]byte(p.Plaintext), 14)
	p.CipherText = string(CipherBytes)
	return err
}
