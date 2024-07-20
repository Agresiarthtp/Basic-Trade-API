package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPass(p string) string {
	cost := 8
	password := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		log.Println("Hashing error:", err)
	}
	return string(hash)
}

func ComparePass(hash, pass []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, pass)
	if err != nil {
		log.Println("Password comparison error:", err)
	}
	return err == nil
}
