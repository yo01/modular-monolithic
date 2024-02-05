package context

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"modular-monolithic/model"
	"modular-monolithic/security/middleware"
	"modular-monolithic/utils"
)

func PageRequestCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pageRequest := parsePageRequest(r.URL.Query())

		ctx := context.WithValue(r.Context(), middleware.PageRequestCtxKey, pageRequest)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func parsePageRequest(query url.Values) *model.PageRequest {
	pageRequest := &model.PageRequest{
		Filters:  make([]map[string]map[string]interface{}, 0),
		Sort:     "",
		Page:     1,
		PerPage:  10,
		Paginate: 1,
		Search:   "",
	}

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "per_page":
			pageRequest.PerPage, _ = strconv.Atoi(queryValue)
		case "page":
			pageRequest.Page, _ = strconv.Atoi(queryValue)
		case "sort":
			pageRequest.Sort = queryValue
		case "filter[]":
			pageRequest.Filters, _ = utils.ValidateAndReturnFilterMap(value)
		case "paginate":
			pageRequest.Paginate, _ = strconv.Atoi(queryValue)
		case "search":
			pageRequest.Search = queryValue
		}
	}

	return pageRequest
}
