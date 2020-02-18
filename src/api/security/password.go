package security

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Hash is hashing process before saving to table
func Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", hashedPassword), nil
}

// VerifyPassword verifying user password and db password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
