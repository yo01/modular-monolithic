package validation

import (
	"errors"
	"modular-monolithic/model"
	"modular-monolithic/security/middleware"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"
)

func ValidateProductAccess(carrier *mcarrier.Carrier) merror.Error {
	// MAIN VARIABLE
	var res merror.Error
	auth := carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)

	if auth.User.Role.Name == "" {
		res = merror.Error{
			Code:  403,
			Error: errors.New("access denied, you don't have role"),
		}
	} else if auth.User.Role.Name != "admin" {
		res = merror.Error{
			Code:  403,
			Error: errors.New("access denied, you don't have an access to this api"),
		}
	}

	return res
}
