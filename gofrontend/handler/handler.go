// Front-end in Go server
// @jeffotoni
// 04/01/2019

package handler

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/jeffotoni/goapiwebserver/gofrontend/pkg/assets"
	"github.com/jeffotoni/goapiwebserver/gofrontend/pkg/util"
)

// Templates
var loginHomeHTML, loginregisterHTML, forgotPasswordHTML string
var loginHomeTpl *template.Template
var loginRegisterTpl *template.Template
var forgotPasswordTpl *template.Template

// be careful with init in multiple packages
// the inits are running on goruntine, it's dangerous,
// it can run after main main func.
func init() {

	// login
	loginHomeHTML = assets.GoMustAssetString("web/templates/login-theme2.html")
	loginHomeTpl = template.Must(template.New("login_view").Parse(loginHomeHTML))

	// register
	loginregisterHTML = assets.GoMustAssetString("web/templates/login-theme2-register.html")
	loginRegisterTpl = template.Must(template.New("login_register_view").Parse(loginregisterHTML))

	// forgot pass
	forgotPasswordHTML = assets.GoMustAssetString("web/templates/forgot-password.html")
	forgotPasswordTpl = template.Must(template.New("login_forgot_pass").Parse(forgotPasswordHTML))
}

// LoginHandler renders the homepage view template
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/css/login-theme2.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(loginHomeHTML),
	}
	render(w, r, loginHomeTpl, "login_view", fullData)
}

// LoginHandlerRegister renders the homepage view template
func LoginHandlerRegister(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(loginregisterHTML),
	}
	render(w, r, loginRegisterTpl, "login_register_view", fullData)
}

// LoginHandlerRegister renders the homepage view template
func ForgotPassHandlerRegister(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(forgotPasswordHTML),
	}
	render(w, r, forgotPasswordTpl, "login_forgot_pass", fullData)
}

/////////// template func
// Render a template, or server error
func render(w http.ResponseWriter, r *http.Request, tpl *template.Template, name string, data interface{}) {
	buf := new(bytes.Buffer)
	if err := tpl.ExecuteTemplate(buf, name, data); err != nil {
		util.Print("\nRender Error" + err.Error() + "\n")
		return
	}
	w.Write(buf.Bytes())
}

// Push the given resource to the client.
func push(w http.ResponseWriter, resource string) {
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push(resource, nil); err == nil {
			return
		}
	}
}
