// Go Api server
// @jeffotoni
// 2019-01-04

package middleware

import (
	"io"
	"net/http"
	"strings"

	"github.com/jeffotoni/goapiwebserver/goapiserver/logg"
	"github.com/jeffotoni/goapiwebserver/goapiserver/pkg/gjwt"
)

func AuthJwt() Adapter {
	s1 := logg.Start()
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if gjwt.CheckJwt(w, r) {
				h.ServeHTTP(w, r)
			} else {
				msgjson := `{"status":"error","message":"error in Jwt!"}`
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusUnauthorized)
				io.WriteString(w, msgjson)
				logg.Show(r.URL.Path, strings.ToUpper(r.Method), "error", s1)
			}
		})
	}
}
