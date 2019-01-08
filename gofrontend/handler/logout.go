// Front-end in Go server
// @jeffotoni
// 2019-01-04

package handler

import (
	"net/http"

	"github.com/jeffotoni/goapiwebserver/gofrontend/pkg/session"
)

// logout session
func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	// set empty
	session.Set(session.NameSession(), "email", "", w, r)
	// destroy
	session.Destroy(session.NameSession(), "email", w, r)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
