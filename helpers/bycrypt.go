package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(p string) string {
	cost := 8
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, cost)
	return string(hash)
}

func ComparePass(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
