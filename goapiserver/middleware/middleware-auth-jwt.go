// Go Api server
// @jeffotoni
// 2019-01-04

package middleware

import "net/http"
import "github.com/jeffotoni/goapiwebserver/goapiserver/pkg/gjwt"

func AuthJwt() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if gjwt.CheckJwt(w, r) {
				h.ServeHTTP(w, r)
			}
		})
	}
}
