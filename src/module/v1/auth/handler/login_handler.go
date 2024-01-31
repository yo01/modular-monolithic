package handler

import (
	"modular-monolithic/module/v1/auth/dto"
	authService "modular-monolithic/module/v1/auth/service"
	userService "modular-monolithic/module/v1/user/service"
	"net/http"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/mhttp"
	"git.motiolabs.com/library/motiolibs/mresponse"
)

type AuthHandler struct {
	Carrier     *mcarrier.Carrier
	AuthService authService.IAuthService
	UserService userService.IUserService
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		req dto.LoginRequest
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
	resp, merr := h.AuthService.SignIn(req)
	if merr.Error != nil {
		mresponse.Failed(w, merr)
		return
	}

	// Return Response
	mresponse.Success(w, "Success", http.StatusOK, resp)
}