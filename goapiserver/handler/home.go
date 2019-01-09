package handler

import (
	"html/template"
	"io"
	"net/http"
	"regexp"
	"strings"

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

		///////////////////////////////////
		// to create routes dynamically
		// with regex for querys use in endpoints
		endpoint, nameregex := EndpointRegex(r)

		//////////////////////////////////////////////////////////////
		// Login regex, search email and validate user
		// creating regex rule for endpoint
		if endpoint == SetEndPoint().PostGetLogin && nameregex != "" {
			if RegexEmail(nameregex) {
				// ok
				LoginExist(nameregex, w, r)
				return
			}
		}
		///////////////////////////// end

		//////////////////////////////////////////////////////////////
		// user regex, search email and validate user
		// creating regex rule for endpoint
		if endpoint == SetEndPoint().PostGetUser && nameregex != "" {
			if RegexEmail(nameregex) {
				// ok
				FindGetUser(nameregex, w, r)
				return
			}
		}
		///////////////////////////// end

		// create url expression, fetch some urls GET
		msg := `{"status":"error","message":"this endpoint does not exist"}`
		w.WriteHeader(http.StatusNotAcceptable)
		io.WriteString(w, msg)
	}
}

// generating endpoint and endpoint nameregex
func EndpointRegex(r *http.Request) (string, string) {
	vetPath := strings.Split(r.URL.Path, "/")
	lastpos := len(vetPath) - 1
	nameregex := vetPath[lastpos]
	endpoint := ""
	for _, v := range vetPath {
		if v != nameregex {
			endpoint += "/" + v
		}
	}
	// endpoint
	endpoint = "/" + strings.Trim(endpoint, "/")
	return endpoint, nameregex
}

// regex email valid
func RegexEmail(nameregex string) bool {
	// regex for my endpoint
	var loginRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`) // Contains email valid
	if loginRegex.MatchString(nameregex) {
		return true
	} else {
		return false
	}
}
