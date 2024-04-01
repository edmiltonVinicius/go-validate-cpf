package utils

import "golang.org/x/crypto/bcrypt"

func Hash(data string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(data), cost)
	return string(bytes), err
}

func CompareHash(original, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(original))
	return err == nil
}
