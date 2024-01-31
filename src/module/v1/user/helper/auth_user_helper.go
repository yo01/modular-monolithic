package helper

import (
	"modular-monolithic/module/v1/user/dto"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func BycryptPassword(req dto.CreateUserRequest) (hashPassword string, err error) {
	hashedPassword, err := Hash(req.Password)
	if err != nil {
		return "", err
	}

	hashPassword = string(hashedPassword)

	return
}
