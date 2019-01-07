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
			ok, msgjson := gjwt.CheckBasic(w, r)
			if ok {

				//h.ServeHTTP(w, r)
				//msg := `{"status":"error","message":"error no authorization!"}`
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusOK)
				io.WriteString(w, msgjson)
				logg.Show(r.URL.Path, strings.ToUpper(r.Method), "success", s1)

			} else {

				//msg := `{"status":"error","message":"error no authorization!"}`
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusUnauthorized)
				io.WriteString(w, msgjson)
				logg.Show(r.URL.Path, strings.ToUpper(r.Method), "error", s1)
			}
		})
	}
}
