package helper

import (
	"modular-monolithic/model"
	"modular-monolithic/module/v1/product/dto"
)

func PrepareToProductsResponse(data []model.Product) (resp []dto.ProductResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = make([]dto.ProductResponse, 0)

	for _, detail := range data {
		// MAIN VARIABLE
		newDetail := new(dto.ProductResponse)

		// SET DATA
		newDetail.ID = detail.ID
		newDetail.Name = detail.Name

		resp = append(resp, *newDetail)
	}

	return
}

func PrepareToDetailProductResponse(data *model.Product) (resp *dto.ProductResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = new(dto.ProductResponse)
	resp.ID = data.ID
	resp.Name = data.Name

	return
}
