package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password []byte) (string) {
	password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return ""
	}
	return string(password)
}