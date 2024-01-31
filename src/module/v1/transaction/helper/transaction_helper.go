package helper

import (
	"modular-monolithic/model"
	"modular-monolithic/module/v1/transaction/dto"
)

func PrepareToTransactionsResponse(data []model.Transaction) (resp []dto.TransactionResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = make([]dto.TransactionResponse, 0)

	for _, detail := range data {
		// MAIN VARIABLE
		newDetail := new(dto.TransactionResponse)

		// SET DATA
		newDetail.ID = detail.ID
		newDetail.Name = detail.Name

		resp = append(resp, *newDetail)
	}

	return
}

func PrepareToDetailTransactionResponse(data *model.Transaction) (resp *dto.TransactionResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = new(dto.TransactionResponse)
	resp.ID = data.ID
	resp.Name = data.Name

	return
}