package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"modular-monolithic/config"
	"modular-monolithic/model"
	"modular-monolithic/module/v1/role/dto"

	"git.motiolabs.com/library/motiolibs/mresponse"
	"git.motiolabs.com/library/motiolibs/mtoken"
)

type ctxKey struct {
	name string
}

var (
	AuthUserCtxKey = &ctxKey{"Auth"}
)

var SECRET_KEY = config.Get().JwtKey
var SECRET_EXPIRED = int32(config.Get().JwtExpired)

func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var authorization = r.Header.Get("authorization")
		token := strings.TrimSpace(strings.Replace(authorization, "Bearer", "", 1))

		json.NewEncoder(w).Encode(r)
		token = strings.TrimSpace(token)

		data, err := mtoken.ValidateJWTToken(token, SECRET_KEY)
		if err.Error != nil {
			mresponse.Failed(w, err)
			return
		}

		//claims
		var claims model.Claims
		claimsBytes, _ := json.Marshal(data)
		json.Unmarshal(claimsBytes, &claims)

		auth := &model.Auth{
			User: model.AuthUser{
				ID:       claims.UserID,
				FullName: claims.FullName,
				Role: dto.RoleResponse{
					ID:   claims.RoleID,
					Name: claims.RoleName,
				},
			},
			Token: token,
		}

		ctx := context.WithValue(r.Context(), AuthUserCtxKey, auth)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
