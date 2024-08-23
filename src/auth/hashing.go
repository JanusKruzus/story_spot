package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const hashSalt = "N3zmAr"

func CreatePasswordHash(password string) string {
	saltedPass := password + hashSalt
	bytes, err := bcrypt.GenerateFromPassword([]byte(saltedPass), 14)
	if err != nil {
		fmt.Println("Hashing error: ", err)
		return ""
	}
	return string(bytes)
}

func ComparePasswordHash(password, hash string) error {
	saltedPass := password + hashSalt
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(saltedPass))
	return err
}
