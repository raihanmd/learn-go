package middleware

import (
	"net/http"
	"restful_api/helper"
	"restful_api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API-Key") != "SECRET" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Message: "UNAUTHORIZED",
		}

		helper.WriteToResBody(w, webResponse)
		return
	}

	middleware.Handler.ServeHTTP(w, r)
}
