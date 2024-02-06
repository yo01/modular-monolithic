package helper

import (
	"modular-monolithic/model"
	roleDTO "modular-monolithic/module/v1/role/dto"
	"modular-monolithic/module/v1/user/dto"

	"github.com/google/uuid"
)

func PrepareToUsersResponse(data []model.User) (resp []dto.UserResponse) {
	// Create a map to store unique CartResponse by id
	resultMap := make(map[uuid.UUID]*dto.UserResponse)

	// Iterate through the input data and group by id
	for _, item := range data {
		newItem := &dto.UserResponse{
			ID:       item.ID,
			FullName: item.FullName,
			Email:    item.Email,
		}

		if item.RoleID != uuid.Nil {
			newItem.Role = new(roleDTO.RoleResponse)
			newItem.Role.ID = item.RoleID
			newItem.Role.Name = *item.RoleName
		}

		resultMap[item.ID] = newItem
	}

	// Convert the map values to a slice
	for _, v := range resultMap {
		resp = append(resp, *v)
	}

	return
}

func PrepareToDetailUserResponse(data []model.User) (resp *dto.UserResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = new(dto.UserResponse)

	if len(data) > 0 {
		resp.ID = data[0].ID
		resp.Email = data[0].Email
		resp.FullName = data[0].FullName
		if data[0].RoleID != uuid.Nil {
			resp.Role = new(roleDTO.RoleResponse)
			resp.Role.ID = data[0].RoleID
			resp.Role.Name = *data[0].RoleName
		}
	}

	return
}

func PrepareToLoginDetailUserResponse(data []model.User) (resp *dto.UserLoginResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = new(dto.UserLoginResponse)

	if len(data) > 0 {
		resp.ID = data[0].ID
		resp.Email = data[0].Email
		resp.FullName = data[0].FullName
		resp.Password = *data[0].Password
		if data[0].RoleID != uuid.Nil {
			resp.Role = new(roleDTO.RoleResponse)
			resp.Role.ID = data[0].RoleID
			resp.Role.Name = *data[0].RoleName
		}
	}

	return
}
