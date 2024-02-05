package handler

import (
	"net/http"

	"modular-monolithic/module/v1/auth/dto"
	authService "modular-monolithic/module/v1/auth/service"
	userService "modular-monolithic/module/v1/user/service"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/mhttp"
	"git.motiolabs.com/library/motiolibs/mresponse"

	"go.uber.org/zap"
)

type AuthHandler struct {
	Carrier     *mcarrier.Carrier
	AuthService authService.IAuthService
	UserService userService.IUserService
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest

	// Validate request
	if merr := mhttp.ValidateRequest(r, &req); merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Initialize carrier
	h.Carrier.Context = r.Context()

	// Invoke service to handle login
	resp, merr := h.AuthService.SignIn(req)
	if merr.Error != nil {
		zap.S().Error(merr.Error)
		mresponse.Failed(w, merr)
		return
	}

	// Return successful response
	mresponse.Success(w, "Success", http.StatusOK, resp)
}
