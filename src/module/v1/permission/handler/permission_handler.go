package handler

import (
	"modular-monolithic/module/v1/permission/dto"
	permissionService "modular-monolithic/module/v1/permission/service"
	"modular-monolithic/utils"
	"net/http"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/mhttp"
	"git.motiolabs.com/library/motiolibs/mresponse"
)

type PermissionHandler struct {
	Carrier           *mcarrier.Carrier
	PermissionService permissionService.IPermissionService
}

func (h *PermissionHandler) List(w http.ResponseWriter, r *http.Request) {
	// Init carrier
	h.Carrier.Context = r.Context()

	// Init Service
	resp, merr := h.PermissionService.List()
	if merr.Error != nil {
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, resp)
}

func (h *PermissionHandler) Detail(w http.ResponseWriter, r *http.Request) {
	// Param
	ID := utils.GetID(r)

	// Init carrier
	h.Carrier.Context = r.Context()

	// Init Service
	resp, merr := h.PermissionService.Detail(ID)
	if merr.Error != nil {
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, resp)
}

func (h *PermissionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		req dto.CreatePermissionRequest
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
	merr = h.PermissionService.Save(req)
	if merr.Error != nil {
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
}

func (h *PermissionHandler) Edit(w http.ResponseWriter, r *http.Request) {
	// Param
	ID := utils.GetID(r)

	var (
		req dto.UpdatePermissionRequest
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
	merr = h.PermissionService.Edit(req, ID)
	if merr.Error != nil {
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
}

func (h *PermissionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// Param
	ID := utils.GetID(r)

	// Init carrier
	h.Carrier.Context = r.Context()

	// Init Service
	merr := h.PermissionService.Delete(ID)
	if merr.Error != nil {
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
}
