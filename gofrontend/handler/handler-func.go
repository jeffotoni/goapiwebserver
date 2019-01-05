// Front-end in Go server
// @jeffotoni
// 04/01/2019

package handler

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/jeffotoni/goapiwebserver/gofrontend/pkg/util"
)

/////////// template func
// Render a template, or server error
func Render(w http.ResponseWriter, r *http.Request, tpl *template.Template, name string, data interface{}) {
	buf := new(bytes.Buffer)
	if err := tpl.ExecuteTemplate(buf, name, data); err != nil {
		util.Print("\nRender Error" + err.Error() + "\n")
		return
	}
	w.Write(buf.Bytes())
}

// Push the given resource to the client.
func Push(w http.ResponseWriter, resource string) {
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push(resource, nil); err == nil {
			return
		}
	}
}
