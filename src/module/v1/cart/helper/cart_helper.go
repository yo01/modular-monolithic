package helper

import (
	"modular-monolithic/model"
	"modular-monolithic/module/v1/cart/dto"
)

func PrepareToCartsResponse(data []model.Cart) (resp []dto.CartResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = make([]dto.CartResponse, 0)

	for _, detail := range data {
		// MAIN VARIABLE
		newDetail := new(dto.CartResponse)

		// SET DATA
		newDetail.ID = detail.ID
		newDetail.ProductID = detail.ProductID

		resp = append(resp, *newDetail)
	}

	return
}

func PrepareToDetailCartResponse(data *model.Cart) (resp *dto.CartResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = new(dto.CartResponse)
	resp.ID = data.ID
	resp.ProductID = data.ProductID

	return
}
