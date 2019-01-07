package handler

import (
	"html/template"
	"io"
	"net/http"

	"github.com/jeffotoni/goapiwebserver/goapiserver/pkg/assets"
)

// Templates
var indexHtml string
var indexTpl *template.Template

// be careful with init in multiple packages
// the inits are running on goruntine, it's dangerous,
// it can run after main main func.
func init() {

	// login
	indexHtml = assets.GoMustAssetString("templates/index.html")
	indexTpl = template.Must(template.New("index_view").Parse(indexHtml))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//accept only "/"
	if r.URL.Path == "/" {

		fullData := map[string]interface{}{
			"NavigationBar": template.HTML(indexHtml),
			"Name":          "Go ApiServer",
		}
		assets.Render(w, r, indexTpl, "index_view", fullData)
	} else {

		msg := `{"status":"error","message":"this endpoint does not exist"}`
		w.WriteHeader(http.StatusNotAcceptable)
		io.WriteString(w, msg)
	}
}
