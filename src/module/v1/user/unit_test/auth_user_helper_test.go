package unit_test

import (
	"testing"

	"modular-monolithic/module/v1/user/dto"
	"modular-monolithic/module/v1/user/helper"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHash(t *testing.T) {
	password := "testPassword"

	hashedPassword, err := helper.Hash(password)
	assert.NoError(t, err)

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	assert.NoError(t, err)
}

func TestBycryptPassword(t *testing.T) {
	req := dto.CreateUserRequest{
		Password: "testPassword",
	}

	hashPassword, err := helper.BycryptPassword(req)
	assert.NoError(t, err)

	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(req.Password))
	assert.NoError(t, err)
}
