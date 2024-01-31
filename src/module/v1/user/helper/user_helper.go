package helper

import (
	"modular-monolithic/model"
	"modular-monolithic/module/v1/user/dto"

	"golang.org/x/crypto/bcrypt"
)

func PrepareToUsersResponse(data []model.User) (resp []dto.UserResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = make([]dto.UserResponse, 0)

	for _, detail := range data {
		// MAIN VARIABLE
		newDetail := new(dto.UserResponse)

		// SET DATA
		newDetail.ID = detail.ID
		newDetail.Email = detail.Email
		newDetail.FullName = detail.FullName

		resp = append(resp, *newDetail)
	}

	return
}

func PrepareToDetailUserResponse(data *model.User) (resp *dto.UserResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = new(dto.UserResponse)
	resp.ID = data.ID
	resp.Email = data.Email
	resp.FullName = data.FullName

	return
}

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
