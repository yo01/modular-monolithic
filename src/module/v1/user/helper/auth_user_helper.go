package helper

import (
	"modular-monolithic/module/v1/user/dto"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func BycryptPassword(req dto.CreateUserRequest) (hashPassword string, err error) {
	hashedPassword, err := Hash(req.Password)
	if err != nil {
		zap.S().Error(err)
		return "", err
	}

	hashPassword = string(hashedPassword)

	return
}
