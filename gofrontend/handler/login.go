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
var loginHomeHTML string
var loginHomeTpl *template.Template

// be careful with init in multiple packages
// the inits are running on goruntine, it's dangerous,
// it can run after main main func.
func init() {

	// login
	loginHomeHTML = assets.GoMustAssetString("templates/login-theme2.html")
	loginHomeTpl = template.Must(template.New("login_view").Parse(loginHomeHTML))
}

// LoginHandler renders the homepage view template
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	//var logado bool
	msgErr := "Sign in"

	if r.Method == http.MethodPost {

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "form error in parse", http.StatusInternalServerError)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" {
			msgErr = "form error username!"
		}

		if password == "" {
			msgErr = "form error password!"
		}

		// make request in apiserver and validate user
		// Do call to API server here

		// if true create session and redirect admin
		// if false not create session
		// sessionName := session.Get("login", username, w, r)

		session.Set("session_user", "username", "jefferson otoni lima", w, r)

		msgErr = "Session Creada"
		//msgErr = "wrong password"

		tplLoginHtml(msgErr, w, r)

		/// redirecionar para /admin se tiver tudo ok
		// logado = true

	} else {

		sessionName := session.Get("session_user", "username", w, r)
		msgErr = "Ola, " + sessionName

		tplLoginHtml(msgErr, w, r)
	}

	//////// HTML SCREEN
	///
	///
	// saving login session
	//session.Set("login", "jeff", "esta logado jefferson", w, r)
	//util.Print("session exist: " + sessionName + "\n")

}

func tplLoginHtml(msgerror string, w http.ResponseWriter, r *http.Request) {

	// if you have logged in enter the admin
	// environment otherwise stay logged in
	assets.Push(w, "/static/css/login-theme2.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(loginHomeHTML),
		"MsgError":      msgerror,
	}

	// ok destroy session
	// session.Destroy("login", "jeff", w, r)
	assets.Render(w, r, loginHomeTpl, "login_view", fullData)
}
