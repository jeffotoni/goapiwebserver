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
	"github.com/jeffotoni/goapiwebserver/gofrontend/pkg/session"
	"github.com/jeffotoni/goapiwebserver/gofrontend/pkg/util"
)

type RegisterUser struct {
	Lastname  string
	Firstname string
	Phone     string
	Email     string
	MsgErr    string
}

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
	var rUserNew = &RegisterUser{}
	rUserNew.MsgErr = msgErr

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
		confirm := r.FormValue("confirm")

		if firstname == "" {
			rUserNew.MsgErr = "form error first name!"
			tplRegisterHtml(rUserNew, w, r)
			return
		} else {
			rUserNew.Firstname = firstname
		}

		if lastname == "" {
			rUserNew.MsgErr = "form error last name!"
			tplRegisterHtml(rUserNew, w, r)
			return
		} else {
			rUserNew.Lastname = lastname
		}

		if email == "" {
			rUserNew.MsgErr = "form error email!"
			tplRegisterHtml(rUserNew, w, r)
			return
		}

		// valid email
		if !util.ValidEmail(email) {
			rUserNew.MsgErr = "form error email invalid!"
			tplRegisterHtml(rUserNew, w, r)
			return
		} else {
			rUserNew.Email = email
		}

		if phone == "" {
			rUserNew.MsgErr = "form error phone!"
			tplRegisterHtml(rUserNew, w, r)
			return
		} else {
			rUserNew.Phone = phone
		}

		if password == "" {
			rUserNew.MsgErr = "form error password!"
			tplRegisterHtml(rUserNew, w, r)
			return
		}

		if confirm == "" {
			rUserNew.MsgErr = "form error confirm!"
			tplRegisterHtml(rUserNew, w, r)
			return
		}

		if confirm != password {
			rUserNew.MsgErr = "Your password is not the same!"
			tplRegisterHtml(rUserNew, w, r)
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
					rUserNew.MsgErr = "User exists try another one!"
					tplRegisterHtml(rUserNew, w, r)
				} else {
					// create new user
					if login.CretaeNew(firstname, lastname, phone, email, password) {
						// new session
						session.Set(session.NameSession(), "email", email, w, r)
						// sign user, create session
						http.Redirect(w, r, "/admin", http.StatusSeeOther)
					} else {
						rUserNew.MsgErr = "something very strange happened with your data, try again!"
						tplRegisterHtml(rUserNew, w, r)
					}
				}
			}
		} else {
			// error
			rUserNew.MsgErr = "something very strange happened, try again!"
			tplRegisterHtml(rUserNew, w, r)
		}

	} else {
		//sessionName := session.Get("session_user", "email", w, r)
		rUserNew.MsgErr = ""
		tplRegisterHtml(rUserNew, w, r)
	}
}

// LoginHandlerRegister renders the homepage view template
func tplRegisterHtml(ruser *RegisterUser, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(loginregisterHTML),
		"msgErr":        ruser.MsgErr,
		"Firstname":     ruser.Firstname,
		"Lastname":      ruser.Lastname,
		"Phone":         ruser.Phone,
		"Email":         ruser.Email,
	}
	assets.Render(w, r, loginRegisterTpl, "login_register_view", fullData)
}
