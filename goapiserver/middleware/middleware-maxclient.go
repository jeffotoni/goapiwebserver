// Go Api server
// @jeffotoni
// 2019-01-04

package handler

import "net/http"

func MaxClients(h http.Handler, n int) http.Handler {

	sema := make(chan struct{}, n)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sema <- struct{}{}

		defer func() { <-sema }()

		h.ServeHTTP(w, r)
	})
}

// func Adapt(h http.Handler, adapters ...Adapter) http.Handler

// func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
//   for _, adapter := range adapters {
//     h = adapter(h)
//   }
//   return h
// }

// http.Handle("/", Adapt(indexHandler, AddHeader("Server", "Mine"),
//                                      CheckAuth(providers),
//                                      CopyMgoSession(db),
//                                      Notify(logger),
//                                    )

// type Adapter func(http.Handler) http.Handler

// func Notify(logger *log.Logger) Adapter {
//   return func(h http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//       logger.Println("before")
//       defer logger.Println("after")
//       h.ServeHTTP(w, r)
//     }
//   }
// }
