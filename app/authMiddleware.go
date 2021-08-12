package app

import (
	domain2 "banking/domain"
	"banking/errs"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

// Port implementation
type AuthMiddleware struct {
	repo domain2.AuthRepository
}

func (a AuthMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r) // &{0xf9d600 false GetAllCustomers <nil> map[GetAllCustomers:0xc000062d20 GetCustomer:0xc000062f00 NewAccount:0xc000063220 NewTransaction:0xc000063540] {false false false {<nil> 0xc0001485b0 []} [0xc0001485b0 [GET]]  <nil>}}
			currentRouteVars := mux.Vars(r)     // ex.: map["customer_id: 2001"]
			authHeader := r.Header.Get("Authorization")

			if authHeader != "" {

				token := getTokenFromHeader(authHeader)

				// checks on authorization rights
				// ex.: GetAllCustomers; GetCustomer (taken from routes in app.go)
				isAuthorized := a.repo.IsAuthorized(token, currentRoute.GetName(), currentRouteVars)

				if isAuthorized {
					next.ServeHTTP(w, r)
				} else {
					appError := errs.AppErr{http.StatusForbidden, "Unauthorized"}
					writeResponse(w, appError.Code, appError.AsMessage())
				}
			} else {
				writeResponse(w, http.StatusUnauthorized, "missing token")
			}
		})
	}
}

func getTokenFromHeader(header string) string {
	/*
	   token is coming in the format as below
	   "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50cyI6W.yI5NTQ3MCIsIjk1NDcyIiw"
	*/
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
