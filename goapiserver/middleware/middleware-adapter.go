// Go Api server
// @jeffotoni
// 2019-01-04

package middleware

import "net/http"

type Adapter func(http.Handler) http.Handler

func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}
