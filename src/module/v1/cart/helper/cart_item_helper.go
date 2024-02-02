package helper

import (
	"modular-monolithic/model"
	"modular-monolithic/module/v1/cart/dto"
)

func PrepareToCartItemsResponse(data []model.CartItem) (resp []dto.CartItemResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = make([]dto.CartItemResponse, 0)

	for _, detail := range data {
		// MAIN VARIABLE
		newDetail := new(dto.CartItemResponse)

		// SET DATA
		newDetail.ID = detail.ID
		newDetail.ProductID = detail.ProductID

		resp = append(resp, *newDetail)
	}

	return
}

func PrepareToDetailCartItemResponse(data *model.CartItem) (resp *dto.CartItemResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = new(dto.CartItemResponse)
	resp.ID = data.ID
	resp.ProductID = data.ProductID

	return
}
