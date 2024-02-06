package handler

import (
	"net/http"

	"modular-monolithic/module/v1/product/dto"
	productService "modular-monolithic/module/v1/product/service"
	"modular-monolithic/utils"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/mhttp"
	"git.motiolabs.com/library/motiolibs/mresponse"

	"go.uber.org/zap"
)

type ProductHandler struct {
	Carrier        *mcarrier.Carrier
	ProductService productService.IProductService
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	// Init carrier
	h.Carrier.Context = r.Context()
	subRouterName := utils.GetSubRouterName(r)

	// Init Service
	resp, merr := h.ProductService.List(subRouterName)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, resp)
}

func (h *ProductHandler) Detail(w http.ResponseWriter, r *http.Request) {
	// Param
	ID := utils.GetID(r)
	subRouterName := utils.GetSubRouterName(r)

	// Init carrier
	h.Carrier.Context = r.Context()

	// Init Service
	resp, merr := h.ProductService.Detail(ID, subRouterName)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, resp)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		req dto.CreateProductRequest
	)
	subRouterName := utils.GetSubRouterName(r)

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
	merr = h.ProductService.Save(req, subRouterName)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
}

func (h *ProductHandler) Edit(w http.ResponseWriter, r *http.Request) {
	// Param
	ID := utils.GetID(r)

	var (
		req dto.UpdateProductRequest
	)
	subRouterName := utils.GetSubRouterName(r)

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
	merr = h.ProductService.Edit(req, ID, subRouterName)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// Param
	ID := utils.GetID(r)
	subRouterName := utils.GetSubRouterName(r)

	// Init carrier
	h.Carrier.Context = r.Context()

	// Init Service
	merr := h.ProductService.Delete(ID, subRouterName)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
}
