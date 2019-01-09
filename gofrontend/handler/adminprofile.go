// Front-end in Go server
// @jeffotoni
// 2019-01-09

package handler

import (
	"html/template"
	"net/http"

	"github.com/jeffotoni/goapiwebserver/gofrontend/api/token"
	"github.com/jeffotoni/goapiwebserver/gofrontend/api/user"
	"github.com/jeffotoni/goapiwebserver/gofrontend/pkg/assets"
	"github.com/jeffotoni/goapiwebserver/gofrontend/pkg/session"
)

type UserProfile struct {
	Lastname  string
	Firstname string
	Phone     string
	Email     string
	MsgErr    string
	Green     string
	Red       string
}

// Templates
var userprofileHTML string
var userprofileTPL *template.Template

// be careful with init in multiple packages
// the inits are running on goruntine, it's dangerous,
// it can run after main main func.
func init() {
	userprofileHTML = assets.GoMustAssetString("templates/admin/admin.html")
	userprofileTPL = template.Must(template.New("admin_index").Parse(userprofileHTML))
}

// LoginHandlerRegister renders the homepage view template
func AdminProfileHandler(w http.ResponseWriter, r *http.Request) {
	msgErr := ""
	var PUser = &UserProfile{}
	PUser.Red = `red`
	PUser.MsgErr = msgErr
	if r.Method == http.MethodPost {
		// fire a post to update the data
		tokenIs := token.FindToken()
		// exist, use
		if tokenIs != "" {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "form error in parse", http.StatusInternalServerError)
				return
			}

			firstname := r.FormValue("firstname")
			lastname := r.FormValue("lastname")
			phone := r.FormValue("phone")

			if firstname == "" {
				PUser.MsgErr = "form error first name!"
				tplUserProfileHtml(PUser, w, r)
				return
			} else {
				PUser.Firstname = firstname
			}

			if lastname == "" {
				PUser.MsgErr = "form error last name!"
				tplUserProfileHtml(PUser, w, r)
				return
			} else {
				PUser.Lastname = lastname
			}

			if phone == "" {
				PUser.MsgErr = "form error phone!"
				tplUserProfileHtml(PUser, w, r)
				return
			} else {
				PUser.Phone = phone
			}

			// id/email only in session
			email := session.Get(session.NameSession(), "email", w, r)
			if email == "" {
				PUser.MsgErr = "error in session!"
				tplUserProfileHtml(PUser, w, r)
			}

			// make request in apiserver and validate user
			// Do call to API server here
			jsonToken := token.FindToken()
			// exist, use
			if jsonToken != "error" {
				// check user and password
				// call api login
				if user.UpdateProfile(jsonToken, firstname, lastname, phone, email) {
					// can not create user
					PUser.MsgErr = "updated successfully!"
					PUser.Green = `green'`
					tplUserProfileHtml(PUser, w, r)
				} else {

					PUser.MsgErr = "error updating in database"
					tplUserProfileHtml(PUser, w, r)

				}
			} else {
				// error
				PUser.MsgErr = "something very strange happened, try again!"
				tplUserProfileHtml(PUser, w, r)
			}

		} else {

			PUser.MsgErr = ""
			tplUserProfileHtml(PUser, w, r)
		}
	} else {
		// Do call to API server here
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}

}

// LoginHandlerRegister renders the homepage view template
func tplUserProfileHtml(ruser *UserProfile, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(userprofileHTML),
		"MsgErr":        ruser.MsgErr,
		"Green":         ruser.Green,
		"Red":           ruser.Red,
		"Firstname":     ruser.Firstname,
		"Lastname":      ruser.Lastname,
		"Phone":         ruser.Phone,
		"Email":         ruser.Email,
	}
	assets.Render(w, r, userprofileTPL, "admin_index", fullData)
}
