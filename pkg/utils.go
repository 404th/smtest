package pkg

import (
	"errors"

	"github.com/404th/smtest/config"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), config.PasswordDefaultCost)
	if err != nil {
		return "", errors.New("failed to hash password: " + err.Error())
	}

	return string(hashedPassword), nil
}

func VerifyPassword(password, hashedPassword string) error {
	if password == "" {
		return errors.New("password cannot be empty")
	}
	if hashedPassword == "" {
		return errors.New("hashed password cannot be empty")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return errors.New("password does not match")
		}
		return errors.New("failed to verify password: " + err.Error())
	}

	return nil
}
