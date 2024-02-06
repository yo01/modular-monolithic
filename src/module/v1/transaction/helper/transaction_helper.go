package helper

import (
	"modular-monolithic/model"
	cartDTO "modular-monolithic/module/v1/cart/dto"
	"modular-monolithic/module/v1/transaction/dto"
	"modular-monolithic/utils"

	"github.com/google/uuid"
)

func PrepareToTransactionsResponse(data []model.Transaction) (resp []dto.TransactionResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = make([]dto.TransactionResponse, 0)

	for _, detail := range data {
		// MAIN VARIABLE
		newDetail := new(dto.TransactionResponse)

		// SET DATA
		newDetail.ID = detail.ID
		if detail.CartID != uuid.Nil {
			newDetail.Cart = new(cartDTO.CartResponse)
			newDetail.Cart.ID = detail.CartID
			newDetail.Cart.UserID = detail.CartUserID
		}
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
	if data.CartID != uuid.Nil {
		resp.Cart = new(cartDTO.CartResponse)
		resp.Cart.ID = data.CartID
		resp.Cart.UserID = data.CartUserID
	}
	resp.IsSuccessPayment = utils.NullBoolToBool(data.IsSuccessPayment)
	resp.PaymentDate = data.PaymentDate

	return
}
