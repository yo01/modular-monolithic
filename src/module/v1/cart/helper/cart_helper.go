package helper

import (
	"modular-monolithic/model"
	"modular-monolithic/module/v1/cart/dto"

	"github.com/google/uuid"
)

func PrepareToCartsResponse(data []model.Cart) (resp []dto.CartResponse) {
	// Create a map to store unique CartResponse by id
	resultMap := make(map[uuid.UUID]*dto.CartResponse)

	// Iterate through the input data and group by id
	for _, item := range data {
		if existingItem, ok := resultMap[item.ID]; ok {
			// Append cart items if the id already exists
			if item.CartItemID != uuid.Nil {
				existingItem.CartItem = append(existingItem.CartItem, dto.CartItemReference{
					ID:          item.CartItemID,
					ProductID:   item.ProductID,
					ProductName: item.ProductName,
				})
			}
		} else {
			// Add new item if id does not exist
			newItem := &dto.CartResponse{
				ID:        item.ID,
				UserID:    item.UserID,
				IsSuccess: item.IsSuccess,
			}

			if item.CartItemID != uuid.Nil {
				newItem.CartItem = append(newItem.CartItem, dto.CartItemReference{
					ID:          item.CartItemID,
					ProductID:   item.ProductID,
					ProductName: item.ProductName,
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

func PrepareToDetailCartResponse(data []model.Cart) (resp *dto.CartResponse) {
	// CONVERT TO RESPONSE STRUCT
	resp = new(dto.CartResponse)

	for _, x := range data {
		resp.ID = x.ID
		resp.UserID = x.UserID
		resp.IsSuccess = x.IsSuccess
		if x.CartItemID != uuid.Nil {
			resp.CartItem = append(resp.CartItem, dto.CartItemReference{
				ID:          x.CartItemID,
				ProductID:   x.ProductID,
				ProductName: x.ProductName,
			})
		}
	}

	return
}
