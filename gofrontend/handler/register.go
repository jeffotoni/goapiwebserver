// Front-end in Go server
// @jeffotoni
// 04/01/2019

package handler

import (
	"html/template"
	"net/http"

	"github.com/jeffotoni/goapiwebserver/gofrontend/pkg/assets"
)

// Templates
var loginregisterHTML string
var loginRegisterTpl *template.Template

// be careful with init in multiple packages
// the inits are running on goruntine, it's dangerous,
// it can run after main main func.
func init() {

	// register
	loginregisterHTML = assets.GoMustAssetString("web/templates/login-theme2-register.html")
	loginRegisterTpl = template.Must(template.New("login_register_view").Parse(loginregisterHTML))
}

// LoginHandlerRegister renders the homepage view template
func LoginHandlerRegister(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(loginregisterHTML),
	}
	Render(w, r, loginRegisterTpl, "login_register_view", fullData)
}
