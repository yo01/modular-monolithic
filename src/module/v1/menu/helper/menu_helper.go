package helper

import (
	"modular-monolithic/model"
	"modular-monolithic/module/v1/menu/dto"
)

func PrepareToMenusResponse(data []model.Menu) (resp []dto.MenuResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = make([]dto.MenuResponse, 0)

	for _, detail := range data {
		// MAIN VARIABLE
		newDetail := new(dto.MenuResponse)

		// SET DATA
		newDetail.ID = detail.ID
		newDetail.Name = detail.Name

		resp = append(resp, *newDetail)
	}

	return
}

func PrepareToDetailMenuResponse(data *model.Menu) (resp *dto.MenuResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = new(dto.MenuResponse)
	resp.ID = data.ID
	resp.Name = data.Name

	return
}
