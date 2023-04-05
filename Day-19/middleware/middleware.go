package middleware

import (
	"context"
	"net/http"
	"strings"

	"Day-19/internal/constants"
)

func isPresent(m string, ms []string) bool {
	for _, v := range ms {
		if m == v {
			return true
		}
	}

	return false
}
func Middle(handler http.Handler) http.Handler {
	vals := make(map[string][]string)
	vals["product-r"] = []string{"GET", "product", "brand"}
	vals["product-w"] = []string{"POST", "product", "brand"}
	vals["brand-r"] = []string{"GET", "brand"}
	vals["brand-w"] = []string{"POST", "brand"}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		head := r.Header.Get("X-API-KEY")
		key, ok := vals[head]

		if !ok {
			http.Error(w, "Unauthorized Request", http.StatusUnauthorized)
			return
		}

		if !isPresent(r.Method, key) {
			http.Error(w, "Forbidden Request", http.StatusForbidden)
			return
		}

		path := strings.Split(r.URL.String(), "/")[1]
		path = strings.Split(path, "?")[0]

		if !isPresent(path, key) {
			http.Error(w, "Forbidden Request", http.StatusForbidden)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func MiddleOrg(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		org := r.Header.Get("organization")
		if r.Method == http.MethodPost {
			ctx := context.WithValue(r.Context(), constants.CtxValue, org)
			r = r.WithContext(ctx)
		}
		handler.ServeHTTP(w, r)
	})
}
