package handler

import (
	"net/http"

	"modular-monolithic/module/v1/role/dto"
	roleService "modular-monolithic/module/v1/role/service"
	"modular-monolithic/utils"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/mhttp"
	"git.motiolabs.com/library/motiolibs/mresponse"

	"go.uber.org/zap"
)

type RoleHandler struct {
	Carrier     *mcarrier.Carrier
	RoleService roleService.IRoleService
}

func (h *RoleHandler) List(w http.ResponseWriter, r *http.Request) {
	// MAIN VARIABLE
	pagination := utils.GeneratePaginationFromRequest(r)

	// Init carrier
	h.Carrier.Context = r.Context()

	// Init Service
	resp, merr := h.RoleService.List(&pagination)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, resp)
}

func (h *RoleHandler) Detail(w http.ResponseWriter, r *http.Request) {
	// Param
	ID := utils.GetID(r)

	// Init carrier
	h.Carrier.Context = r.Context()

	// Init Service
	resp, merr := h.RoleService.Detail(ID)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, resp)
}

func (h *RoleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		req dto.CreateRoleRequest
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
	merr = h.RoleService.Save(req)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
}

func (h *RoleHandler) Edit(w http.ResponseWriter, r *http.Request) {
	// Param
	ID := utils.GetID(r)

	var (
		req dto.UpdateRoleRequest
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
	merr = h.RoleService.Edit(req, ID)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
}

func (h *RoleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// Param
	ID := utils.GetID(r)

	// Init carrier
	h.Carrier.Context = r.Context()

	// Init Service
	merr := h.RoleService.Delete(ID)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
}
