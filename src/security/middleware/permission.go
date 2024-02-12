package middleware

import (
	"errors"
	"net/http"

	"modular-monolithic/constant"
	"modular-monolithic/model"
	"modular-monolithic/utils"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"go.uber.org/zap"
)

func ValidateAccess(carrier *mcarrier.Carrier, accessType, subRouterName string, conditions []string) merror.Error {
	// MAIN VARIABLE
	var res merror.Error

	switch accessType {
	case constant.AccessTypePermission:
		res = ValidatePermissionAccess(carrier)
	default:
		res = GeneralValidationAccess(carrier, subRouterName, conditions)
	}

	return res
}

func GeneralValidationAccess(carrier *mcarrier.Carrier, subRouterName string, conditions []string) merror.Error {
	// MAIN VARIABLE
	var res merror.Error
	var context = carrier.Context.Value(AuthUserCtxKey)

	if context != nil {
		auth := context.(*model.Auth)

		if auth.User.Role.Name == "" {
			err := errors.New("access denied, you don't have role")
			zap.S().Error(err)
			res = merror.Error{
				Code:  http.StatusForbidden,
				Error: err,
			}
		} else if auth.User.Role.Name != "admin" {
			err := errors.New("access denied, only admin can access this api")
			zap.S().Error(err)
			res = merror.Error{
				Code:  http.StatusForbidden,
				Error: err,
			}
		} else if auth.User.Role.Name == "admin" {
			if !utils.InStringArray(subRouterName, conditions) {
				err := errors.New("access denied, you don't have an access to this api")
				zap.S().Error(err)
				res = merror.Error{
					Code:  http.StatusForbidden,
					Error: err,
				}
			}
		}
	}

	return res
}

func ValidatePermissionAccess(carrier *mcarrier.Carrier) merror.Error {
	// MAIN VARIABLE
	var res merror.Error
	var context = carrier.Context.Value(AuthUserCtxKey)

	if context != nil {
		auth := context.(*model.Auth)

		if auth.User.Role.Name == "" {
			err := errors.New("access denied, you don't have role")
			zap.S().Error(err.Error)
			res = merror.Error{
				Code:  http.StatusForbidden,
				Error: err,
			}
		} else if auth.User.Role.Name != "admin" {
			err := errors.New("access denied, you don't have an access to this api")
			zap.S().Error(err.Error)
			res = merror.Error{
				Code:  http.StatusForbidden,
				Error: err,
			}
		}
	}

	return res
}
