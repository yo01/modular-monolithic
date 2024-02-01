package unit_test

import (
	"context"
	"errors"

	"testing"

	"modular-monolithic/model"
	"modular-monolithic/module/v1/role/dto"
	"modular-monolithic/module/v1/user/validation"
	"modular-monolithic/security/middleware"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"
)

func TestValidateMenuAccess(t *testing.T) {
	// Mocking the context value for an admin user
	adminAuth := &model.Auth{
		User: model.AuthUser{
			Role: dto.RoleResponse{
				Name: "admin",
			},
		},
	}

	// Mocking the context value for a user without a role
	otherAuth := &model.Auth{
		User: model.AuthUser{
			Role: dto.RoleResponse{
				Name: "member",
			},
		},
	}

	tests := []struct {
		name           string
		auth           *model.Auth
		expectedResult merror.Error
	}{
		{
			name:           "Admin user",
			auth:           adminAuth,
			expectedResult: merror.Error{},
		},
		{
			name: "Other user",
			auth: otherAuth,
			expectedResult: merror.Error{
				Code:  403,
				Error: errors.New("access denied, you don't have an access to this api"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Creating a mock context
			mockContext := &mcarrier.Carrier{
				Context: context.WithValue(context.Background(), middleware.AuthUserCtxKey, tt.auth),
			}

			// Calling the function under test
			result := validation.ValidateUserAccess(mockContext)

			// Asserting the result
			if result.Code != tt.expectedResult.Code && result.Error != tt.expectedResult.Error {
				t.Errorf("Test case %s failed. Expected %v, got %v", tt.name, tt.expectedResult, result)
			}
		})
	}
}
