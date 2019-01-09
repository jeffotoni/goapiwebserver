// Front-end in Go server
// @jeffotoni
// 2019-01-04

package handler

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/jeffotoni/goapiwebserver/gofrontend/api/token"
	"github.com/jeffotoni/goapiwebserver/gofrontend/api/user"
	"github.com/jeffotoni/goapiwebserver/gofrontend/pkg/assets"
	"github.com/jeffotoni/goapiwebserver/gofrontend/pkg/session"
)

// Templates
var adminHTML string
var adminTPL *template.Template

// be careful with init in multiple packages
// the inits are running on goruntine, it's dangerous,
// it can run after main main func.
func init() {

	// forgot pass
	adminHTML = assets.GoMustAssetString("templates/admin/admin.html")
	adminTPL = template.Must(template.New("admin_index").Parse(adminHTML))
}

// LoginHandlerRegister renders the homepage view template
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	sessionName := session.Get(session.NameSession(), "email", w, r)
	// session user
	if sessionName != "" {
		tokenIs := token.FindToken()
		// exist, token
		if tokenIs != "" {
			// search the user data and play on the screen
			jsonUser := user.SearchUser(tokenIs, sessionName)
			if jsonUser != "error" {
				type UserJ struct {
					Firstname string
					Lastname  string
					Phone     string
					Email     string
				}
				var U = UserJ{}
				var jsonByte = []byte(jsonUser)
				err := json.Unmarshal(jsonByte, &U)
				if err == nil {
					// fill in the form fields
					fullData := map[string]interface{}{
						"NavigationBar": template.HTML(adminHTML),
						"Firstname":     U.Firstname,
						"Lastname":      U.Lastname,
						"Phone":         U.Phone,
						"Email":         U.Email,
					}
					assets.Render(w, r, adminTPL, "admin_index", fullData)
					return
				}
			}
		}

		fullData := map[string]interface{}{
			"NavigationBar": template.HTML(adminHTML),
		}
		assets.Render(w, r, adminTPL, "admin_index", fullData)

	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
