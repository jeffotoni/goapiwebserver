// Front-end in Go server
// @jeffotoni
// 2019-01-04

package handler

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/jeffotoni/goapiwebserver/gofrontend/api/login"
	"github.com/jeffotoni/goapiwebserver/gofrontend/api/token"
	"github.com/jeffotoni/goapiwebserver/gofrontend/pkg/assets"
	"github.com/jeffotoni/goapiwebserver/gofrontend/pkg/util"
)

// Templates
var loginregisterHTML string
var loginRegisterTpl *template.Template

// be careful with init in multiple packages
// the inits are running on goruntine, it's dangerous,
// it can run after main main func.
func init() {

	// register
	loginregisterHTML = assets.GoMustAssetString("templates/login-theme2-register.html")
	loginRegisterTpl = template.Must(template.New("login_register_view").Parse(loginregisterHTML))
}

// LoginHandler renders the homepage view template
func LoginHandlerRegister(w http.ResponseWriter, r *http.Request) {

	//var bool
	msgErr := ""

	if r.Method == http.MethodPost {

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "form error in parse", http.StatusInternalServerError)
			return
		}

		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		email := r.FormValue("email")
		phone := r.FormValue("phone")
		password := r.FormValue("password")
		confirm := r.FormValue("confifm")

		if firstname == "" {
			msgErr = "form error first name!"
			tplRegisterHtml(msgErr, w, r)
			return
		}

		if lastname == "" {
			msgErr = "form error last name!"
			tplRegisterHtml(msgErr, w, r)
			return
		}

		if email == "" {
			msgErr = "form error email!"
			tplRegisterHtml(msgErr, w, r)
			return
		}

		// valid email
		if !util.ValidEmail(email) {
			msgErr = "form error email invalid!"
			tplRegisterHtml(msgErr, w, r)
			return
		}

		if phone == "" {
			msgErr = "form error phone!"
			tplRegisterHtml(msgErr, w, r)
			return
		}

		if password == "" {
			msgErr = "form error password!"
			tplRegisterHtml(msgErr, w, r)
			return
		}

		if confirm == "" {
			msgErr = "form error confirm!"
			tplRegisterHtml(msgErr, w, r)
			return
		}

		if confirm != password {
			msgErr = "Your password is not the same!"
			tplRegisterHtml(msgErr, w, r)
			return
		}

		// make request in apiserver and validate user
		// Do call to API server here
		jsonToken := token.GetTokenApiServer()
		// exist, use
		if jsonToken != "error" {
			//struct message token server
			var stToken = &token.TokenStruct{}
			json.Unmarshal([]byte(jsonToken), &stToken)
			// token exists, can continue
			if stToken.Token != "" && len(stToken.Expires) == 10 {
				// check user and password
				// call api login
				if login.ExistUser(stToken.Token, email) == "success" {
					// can not create user
					msgErr = "User exists try another one!"
					tplRegisterHtml(msgErr, w, r)
				} else {

					// create new user
					if login.CretaeNew(firstname, lastname, phone, email, password) {
						http.Redirect(w, r, "/register-success", http.StatusSeeOther)
					} else {
						msgErr = "something very strange happened with your data, try again!"
						tplRegisterHtml(msgErr, w, r)
					}
				}
			}
		} else {
			// error
			msgErr = "something very strange happened, try again!"
			tplRegisterHtml(msgErr, w, r)
		}

	} else {
		//sessionName := session.Get("session_user", "email", w, r)
		msgErr = ""
		tplRegisterHtml(msgErr, w, r)
	}
}

// LoginHandlerRegister renders the homepage view template
func tplRegisterHtml(msgErr string, w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(loginregisterHTML),
		"msgErr":        msgErr,
	}
	assets.Render(w, r, loginRegisterTpl, "login_register_view", fullData)
}
