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
var forgotPasswordHTML string
var forgotPasswordTpl *template.Template

// be careful with init in multiple packages
// the inits are running on goruntine, it's dangerous,
// it can run after main main func.
func init() {

	// forgot pass
	forgotPasswordHTML = assets.GoMustAssetString("web/templates/forgot-password.html")
	forgotPasswordTpl = template.Must(template.New("login_forgot_pass").Parse(forgotPasswordHTML))
}

// LoginHandlerRegister renders the homepage view template
func ForgotPassHandlerRegister(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(forgotPasswordHTML),
	}
	Render(w, r, forgotPasswordTpl, "login_forgot_pass", fullData)
}
