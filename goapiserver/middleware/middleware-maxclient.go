// Go Api server
// @jeffotoni
// 2019-01-04

package middleware

import (
	"net/http"
)

func MaxClients(n int) Adapter {
	sema := make(chan struct{}, n)
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sema <- struct{}{}
			defer func() { <-sema }()
			h.ServeHTTP(w, r)
		})
	}
}
