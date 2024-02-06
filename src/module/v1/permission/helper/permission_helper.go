package helper

import (
	"modular-monolithic/model"
	"modular-monolithic/module/v1/permission/dto"
	roleDTO "modular-monolithic/module/v1/role/dto"

	"github.com/google/uuid"
)

func PrepareToPermissionsResponse(data []model.Permission) (resp []dto.PermissionResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = make([]dto.PermissionResponse, 0)

	for _, detail := range data {
		// MAIN VARIABLE
		newDetail := new(dto.PermissionResponse)

		// SET DATA
		newDetail.ID = detail.ID
		newDetail.ListAPI = detail.ListAPI
		if detail.RoleID != uuid.Nil {
			newDetail.Role = new(roleDTO.RoleResponse)
			newDetail.Role.ID = detail.RoleID
			newDetail.Role.Name = *detail.RoleName
		}

		resp = append(resp, *newDetail)
	}

	return
}

func PrepareToDetailPermissionResponse(data *model.Permission) (resp *dto.PermissionResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = new(dto.PermissionResponse)
	resp.ID = data.ID
	resp.ListAPI = data.ListAPI
	if data.RoleID != uuid.Nil {
		resp.Role = new(roleDTO.RoleResponse)
		resp.Role.ID = data.RoleID
		resp.Role.Name = *data.RoleName
	}

	return
}
