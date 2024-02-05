package helper

import (
	"modular-monolithic/model"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func BycryptPassword(data *model.User) error {
	hashedPassword, err := Hash(*data.Password)
	if err != nil {
		zap.S().Error(err)
		return err
	}

	*data.Password = string(hashedPassword)

	return nil
}
