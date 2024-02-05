package handler

import (
	"fmt"
	"net/http"

	"modular-monolithic/module/v1/cart/dto"
	cartService "modular-monolithic/module/v1/cart/service"
	"modular-monolithic/utils"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"
	"git.motiolabs.com/library/motiolibs/mhttp"
	"git.motiolabs.com/library/motiolibs/mresponse"

	"go.uber.org/zap"

	"github.com/google/uuid"
)

type CartHandler struct {
	Carrier         *mcarrier.Carrier
	CartService     cartService.ICartService
	CartItemService cartService.ICartItemService
}

func (h *CartHandler) List(w http.ResponseWriter, r *http.Request) {
	// Init carrier
	h.Carrier.Context = r.Context()

	// Init Service
	resp, merr := h.CartService.List()
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, resp)
}

func (h *CartHandler) Detail(w http.ResponseWriter, r *http.Request) {
	// Param
	ID := utils.GetID(r)

	// Init carrier
	h.Carrier.Context = r.Context()

	// Init Service
	resp, merr := h.CartService.Detail(ID)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	} else if resp.ID == uuid.Nil {
		err := fmt.Errorf("cart with id %v is not found", ID)
		zap.S().Error(err)
		mresponse.Failed(w, merror.Error{
			Code:  404,
			Error: err,
		})
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, resp)
}

func (h *CartHandler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		req dto.CreateCartRequest
	)

	// Validation
	merr := mhttp.ValidateRequest(r, &req)
	if merr.Error != nil {
		mresponse.Failed(w, merr)
		return
	}

	// Init carrier
	h.Carrier.Context = r.Context()

	// Init Service
	resp, merr := h.CartService.Save(req)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// CART ITEM RELATION (NEED REFACTOR)
	for _, productID := range req.ProductID {
		// TEMP STRUCT
		cartItemData := dto.CreateCartItemRequest{
			CartID:    resp.ID.String(),
			ProductID: productID,
		}

		merr = h.CartItemService.Save(cartItemData)
		if merr.Error != nil {
			zap.S().Error(merr.Error)
			mresponse.Failed(w, merr)
			return
		}
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
}

func (h *CartHandler) Edit(w http.ResponseWriter, r *http.Request) {
	// Param
	ID := utils.GetID(r)

	var (
		req dto.UpdateCartRequest
	)

	// Validation
	merr := mhttp.ValidateRequest(r, &req)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Init carrier
	h.Carrier.Context = r.Context()

	// Init Service
	merr = h.CartService.Edit(req, ID)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// TEMP STRUCT
	cartItemData := dto.UpdateCartItemRequest{
		ProductID: req.ProductID,
	}

	// UPDATE CART ITEM
	merr = h.CartItemService.Edit(cartItemData, req.CartItemID, ID)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
}

func (h *CartHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// Param
	ID := utils.GetID(r)

	// Init carrier
	h.Carrier.Context = r.Context()

	// Init Service
	merr := h.CartService.Delete(ID)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Init Service
	merr = h.CartItemService.Delete(ID)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
}
