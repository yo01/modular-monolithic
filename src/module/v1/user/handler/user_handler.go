package handler

import (
	"net/http"

	"modular-monolithic/module/v1/user/dto"
	userService "modular-monolithic/module/v1/user/service"
	"modular-monolithic/utils"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/mhttp"
	"git.motiolabs.com/library/motiolibs/mresponse"

	"go.uber.org/zap"
)

type UserHandler struct {
	Carrier     *mcarrier.Carrier
	UserService userService.IUserService
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	// Init carrier
	h.Carrier.Context = r.Context()
	subRouterName := utils.GetSubRouterName(r)

	// Init Service
	resp, merr := h.UserService.List(subRouterName)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, resp)
}

func (h *UserHandler) Detail(w http.ResponseWriter, r *http.Request) {
	// Param
	ID := utils.GetID(r)
	subRouterName := utils.GetSubRouterName(r)

	// Init carrier
	h.Carrier.Context = r.Context()

	// Init Service
	resp, merr := h.UserService.Detail(ID, subRouterName)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, resp)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		req dto.CreateUserRequest
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
	merr = h.UserService.Save(req, subRouterName)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
}

func (h *UserHandler) Edit(w http.ResponseWriter, r *http.Request) {
	// Param
	ID := utils.GetID(r)
	subRouterName := utils.GetSubRouterName(r)

	var (
		req dto.UpdateUserRequest
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
	merr = h.UserService.Edit(req, ID, subRouterName)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// Param
	ID := utils.GetID(r)
	subRouterName := utils.GetSubRouterName(r)

	// Init carrier
	h.Carrier.Context = r.Context()

	// Init Service
	merr := h.UserService.Delete(ID, subRouterName)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
}
