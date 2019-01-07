// Go Api server
// @jeffotoni
// 2019-01-04

package middleware

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/didip/tollbooth"
	"github.com/jeffotoni/goapiwebserver/goapiserver/config"
	"github.com/jeffotoni/goapiwebserver/goapiserver/logg"
)

// Creating limiter for all handlers
// or one for each handler. Your choice.
// This limiter basically says: allow at most NewLimiter request per 1 second.
//limiter := tollbooth.NewLimiter(1, nil)
// Limit only GET and POST requests.
//limiter.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"}).SetMethods([]string{"GET", "POST"})
func Limit() Adapter {
	s1 := logg.Start()
	return func(h http.Handler) http.Handler {

		lmt := tollbooth.NewLimiter(config.REQUEST_SEC, nil)
		//lmt = tollbooth.NewLimiter(1, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
		lmt.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})
		lmt.SetMethods([]string{"GET", "POST"})

		// Set a custom expiration TTL for token bucket.
		lmt.SetTokenBucketExpirationTTL(time.Hour)

		// Set a custom expiration TTL for basic auth users.
		lmt.SetBasicAuthExpirationTTL(time.Hour)

		// Set a custom expiration TTL for header entries.
		lmt.SetHeaderEntryExpirationTTL(time.Hour)

		//lmt := tollbooth.NewLimiter(1, nil)

		// Set a custom message.
		lmt.SetMessage("You have reached maximum request limit.")

		// Set a custom content-type.
		lmt.SetMessageContentType("text/plain; charset=utf-8")

		// Set a custom function for rejection.
		lmt.SetOnLimitReached(func(w http.ResponseWriter, r *http.Request) {
			msg := `{"status":"error","message":"error no authorization!"}`
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, msg)
			logg.Show(r.URL.Path, strings.ToUpper(r.Method), "error", s1)
		})
		return tollbooth.LimitFuncHandler(tollbooth.NewLimiter(config.REQUEST_SEC, nil), h.ServeHTTP)
	}
}
