package helper

import (
	"modular-monolithic/model"
	"modular-monolithic/module/v1/transaction/dto"
	"modular-monolithic/utils"
)

func PrepareToTransactionsResponse(data []model.Transaction) (resp []dto.TransactionResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = make([]dto.TransactionResponse, 0)

	for _, detail := range data {
		// MAIN VARIABLE
		newDetail := new(dto.TransactionResponse)

		// SET DATA
		newDetail.ID = detail.ID
		newDetail.CartID = detail.CartID
		newDetail.IsSuccessPayment = utils.NullBoolToBool(detail.IsSuccessPayment)
		newDetail.PaymentDate = detail.PaymentDate

		resp = append(resp, *newDetail)
	}

	return
}

func PrepareToDetailTransactionResponse(data *model.Transaction) (resp *dto.TransactionResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = new(dto.TransactionResponse)
	resp.ID = data.ID
	resp.CartID = data.CartID
	resp.IsSuccessPayment = utils.NullBoolToBool(data.IsSuccessPayment)
	resp.PaymentDate = data.PaymentDate
	resp.PaymentByID = data.PaymentByID

	return
}
