package handler

import (
	"net/http"

	"modular-monolithic/module/v1/user/dto"
	userService "modular-monolithic/module/v1/user/service"
	"modular-monolithic/utils"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/mhttp"
	"git.motiolabs.com/library/motiolibs/mresponse"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	Carrier     *mcarrier.Carrier
	UserService userService.IUserService
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	// Init carrier
	h.Carrier.Context = r.Context()

	// Init Service
	resp, merr := h.UserService.List()
	if merr.Error != nil {
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, resp)
	return
}

func (h *UserHandler) Detail(w http.ResponseWriter, r *http.Request) {
	// Param
	ID := utils.GetID(r)

	// Init carrier
	h.Carrier.Context = r.Context()

	// Init Service
	resp, merr := h.UserService.Detail(ID)
	if merr.Error != nil {
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, resp)
	return
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		req dto.CreateUserRequest
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
	merr = h.UserService.Save(req)
	if merr.Error != nil {
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
	return
}

func (h *UserHandler) Edit(w http.ResponseWriter, r *http.Request) {
	// Param
	vars := mux.Vars(r)
	id := vars["id"]

	var (
		req dto.UpdateUserRequest
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
	merr = h.UserService.Edit(req, id)
	if merr.Error != nil {
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
	return
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// Param
	vars := mux.Vars(r)
	id := vars["id"]

	// Init carrier
	h.Carrier.Context = r.Context()

	// Init Service
	merr := h.UserService.Delete(id)
	if merr.Error != nil {
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, true)
	return
}
