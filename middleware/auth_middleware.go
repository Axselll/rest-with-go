package middleware

import (
	"go-rest/entity/web"
	"go-rest/helper"
	"net/http"
	"os"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (m AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if os.Getenv("APIKey") == r.Header.Get("X-API-KEY") {
		m.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webRes := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResBody(w, webRes)
	}
}
