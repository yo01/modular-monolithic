package handler

import (
	"net/http"

	"modular-monolithic/module/v1/menu/dto"
	menuService "modular-monolithic/module/v1/menu/service"
	"modular-monolithic/utils"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/mhttp"
	"git.motiolabs.com/library/motiolibs/mresponse"

	"go.uber.org/zap"
)

type MenuHandler struct {
	Carrier     *mcarrier.Carrier
	MenuService menuService.IMenuService
}

func (h *MenuHandler) List(w http.ResponseWriter, r *http.Request) {
	// Init carrier
	h.Carrier.Context = r.Context()
	subRouterName := utils.GetSubRouterName(r)

	// Init Service
	resp, merr := h.MenuService.List(subRouterName)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, resp)
}

func (h *MenuHandler) Detail(w http.ResponseWriter, r *http.Request) {
	// Param
	ID := utils.GetID(r)

	// Init carrier
	h.Carrier.Context = r.Context()
	subRouterName := utils.GetSubRouterName(r)

	// Init Service
	resp, merr := h.MenuService.Detail(ID, subRouterName)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, resp)
}

func (h *MenuHandler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		req dto.CreateMenuRequest
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
	merr = h.MenuService.Save(req, subRouterName)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
}

func (h *MenuHandler) Edit(w http.ResponseWriter, r *http.Request) {
	// Param
	ID := utils.GetID(r)

	var (
		req dto.UpdateMenuRequest
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
	merr = h.MenuService.Edit(req, ID, subRouterName)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
}

func (h *MenuHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// Param
	ID := utils.GetID(r)
	subRouterName := utils.GetSubRouterName(r)

	// Init carrier
	h.Carrier.Context = r.Context()

	// Init Service
	merr := h.MenuService.Delete(ID, subRouterName)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
}
