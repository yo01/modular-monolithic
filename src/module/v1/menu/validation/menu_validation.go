package validation

import (
	"errors"

	"modular-monolithic/model"
	"modular-monolithic/security/middleware"
	"modular-monolithic/utils"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"go.uber.org/zap"
)

func ValidateMenuAccess(carrier *mcarrier.Carrier, subRouterName string, conditions []string) merror.Error {
	// MAIN VARIABLE
	var res merror.Error
	var context = carrier.Context.Value(middleware.AuthUserCtxKey)

	if context != nil {
		auth := context.(*model.Auth)

		if auth.User.Role.Name == "" {
			err := errors.New("access denied, you don't have role")
			zap.S().Error(err)
			res = merror.Error{
				Code:  403,
				Error: err,
			}
		} else if auth.User.Role.Name != "admin" {
			err := errors.New("access denied, you don't have an access to this api")
			zap.S().Error(err)
			res = merror.Error{
				Code:  403,
				Error: err,
			}
		} else if auth.User.Role.Name == "admin" {
			if !utils.InStringArray(subRouterName, conditions) {
				err := errors.New("access denied, you don't have an access to this api")
				zap.S().Error(err)
				res = merror.Error{
					Code:  403,
					Error: err,
				}
			}
		}
	}

	return res
}
