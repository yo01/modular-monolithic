package helper

import (
	"modular-monolithic/model"
	roleDTO "modular-monolithic/module/v1/role/dto"
	"modular-monolithic/module/v1/user/dto"
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
		if detail.RoleName != nil {
			newDetail.Role = new(roleDTO.RoleResponse)
			newDetail.Role.ID = detail.RoleID
			newDetail.Role.Name = *detail.RoleName
		}

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
	if data.RoleName != nil {
		resp.Role = new(roleDTO.RoleResponse)
		resp.Role.ID = data.RoleID
		resp.Role.Name = *data.RoleName
	}

	return
}
