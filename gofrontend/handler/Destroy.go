// Front-end in Go server
// @jeffotoni
// 2019-01-04

package handler

import (
	"html/template"
	"net/http"

	"github.com/jeffotoni/goapiwebserver/gofrontend/pkg/assets"
	"github.com/jeffotoni/goapiwebserver/gofrontend/pkg/session"
)

// Templates
var destroyHTML, pagetpl string
var destroyTPL *template.Template

// be careful with init in multiple packages
// the inits are running on goruntine, it's dangerous,
// it can run after main main func.
func init() {

	pagetpl = "destroy"
	// forgot pass
	destroyHTML = assets.GoMustAssetString("templates/admin/destroy.html")
	destroyTPL = template.Must(template.New(pagetpl).Parse(destroyHTML))
}

//
func DestroyHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	sessionName := session.Get("session_user", "username", w, r)
	menssage := "Sua session foi destruida: " + sessionName

	// destroy session
	session.Destroy("session_user", "username", w, r)

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(destroyHTML),
		"Username":      menssage,
	}
	assets.Render(w, r, destroyTPL, pagetpl, fullData)
}
