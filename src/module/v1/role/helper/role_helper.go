package helper

import (
	"modular-monolithic/model"
	"modular-monolithic/module/v1/role/dto"
)

func PrepareToRolesResponse(data []model.Role) (resp []dto.RoleResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = make([]dto.RoleResponse, 0)

	for _, detail := range data {
		// MAIN VARIABLE
		newDetail := new(dto.RoleResponse)

		// SET DATA
		newDetail.ID = detail.ID
		newDetail.Name = detail.Name

		resp = append(resp, *newDetail)
	}

	return
}

func PrepareToDetailRoleResponse(data *model.Role) (resp *dto.RoleResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = new(dto.RoleResponse)
	resp.ID = data.ID
	resp.Name = data.Name

	return
}
