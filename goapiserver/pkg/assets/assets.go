// Back-End in Go server
// @jeffotoni
// 2019-01-04

package assets

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/jeffotoni/goapiwebserver/goapiserver/pkg/util"
)

// GoMustAssetString returns the asset
func GoMustAssetString(strtext string) string {
	return string(MustAsset(strtext))
}

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
