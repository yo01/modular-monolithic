package helper

import (
	"modular-monolithic/model"
	cartDTO "modular-monolithic/module/v1/cart/dto"
	"modular-monolithic/module/v1/transaction/dto"

	"github.com/google/uuid"
)

func PrepareToTransactionsResponse(data []model.Transaction) (resp []dto.TransactionResponse) {
	// Create a map to store unique TransactionResponse by id
	resultMap := make(map[uuid.UUID]*dto.TransactionResponse)

	// Iterate through the input data and group by id
	for _, item := range data {
		if existingItem, ok := resultMap[item.ID]; ok {
			// Append cart items if the id already exists
			if item.CartItemID != uuid.Nil {
				existingItem.Cart.CartItem = append(existingItem.Cart.CartItem, cartDTO.CartItemReference{
					ID:          item.CartItemID,
					ProductID:   item.CartItemProductID,
					ProductName: &item.ProductName,
				})
			}
		} else {
			// Add new item if id does not exist
			newItem := &dto.TransactionResponse{
				ID:               item.ID,
				IsSuccessPayment: item.IsSuccessPayment,
				PaymentDate:      item.PaymentDate,
			}

			if item.CartID != uuid.Nil {
				newItem.Cart = new(cartDTO.CartResponse)
				newItem.Cart.ID = item.CartID
				newItem.Cart.IsSuccess = item.CartIsSuccess
				newItem.Cart.UserID = item.CartUserID
			}

			if item.CartItemID != uuid.Nil {
				newItem.Cart.CartItem = make([]cartDTO.CartItemReference, 0)
				newItem.Cart.CartItem = append(newItem.Cart.CartItem, cartDTO.CartItemReference{
					ID:          item.CartItemID,
					ProductID:   item.CartItemProductID,
					ProductName: &item.ProductName,
				})
			}

			resultMap[item.ID] = newItem
		}
	}

	// Convert the map values to a slice
	for _, v := range resultMap {
		resp = append(resp, *v)
	}

	return
}

func PrepareToDetailTransactionResponse(data []model.Transaction) (resp *dto.TransactionResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = new(dto.TransactionResponse)

	if len(data) > 0 {
		resp.ID = data[0].ID
		resp.PaymentDate = data[0].PaymentDate
		resp.IsSuccessPayment = data[0].IsSuccessPayment
		if data[0].CartID != uuid.Nil {
			resp.Cart = new(cartDTO.CartResponse)
			resp.Cart.ID = data[0].CartID
			resp.Cart.IsSuccess = data[0].CartIsSuccess
			resp.Cart.UserID = data[0].CartUserID
		}
	}

	for _, x := range data {
		if x.CartItemID != uuid.Nil {
			resp.Cart.CartItem = append(resp.Cart.CartItem, cartDTO.CartItemReference{
				ID:          x.CartItemID,
				ProductID:   x.CartItemProductID,
				ProductName: &x.ProductName,
			})
		}
	}

	return
}
