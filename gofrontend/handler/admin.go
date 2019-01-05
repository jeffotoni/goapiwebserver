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
var adminHTML string
var adminTPL *template.Template

// be careful with init in multiple packages
// the inits are running on goruntine, it's dangerous,
// it can run after main main func.
func init() {

	// forgot pass
	adminHTML = assets.GoMustAssetString("web/templates/admin/admin.html")
	adminTPL = template.Must(template.New("admin_index").Parse(adminHTML))
}

// LoginHandlerRegister renders the homepage view template
func AdminHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(adminHTML),
	}
	Render(w, r, adminTPL, "admin_index", fullData)
}
