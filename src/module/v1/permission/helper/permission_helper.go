package helper

import (
	"modular-monolithic/model"
	"modular-monolithic/module/v1/permission/dto"
)

func PrepareToPermissionsResponse(data []model.Permission) (resp []dto.PermissionResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = make([]dto.PermissionResponse, 0)

	for _, detail := range data {
		// MAIN VARIABLE
		newDetail := new(dto.PermissionResponse)

		// SET DATA
		newDetail.ID = detail.ID
		newDetail.Name = detail.Name

		resp = append(resp, *newDetail)
	}

	return
}

func PrepareToDetailPermissionResponse(data *model.Permission) (resp *dto.PermissionResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = new(dto.PermissionResponse)
	resp.ID = data.ID
	resp.Name = data.Name

	return
}
