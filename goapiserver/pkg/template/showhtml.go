// Go Api server
// @jeffotoni
// 2019-01-04

// template
package tpl

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/jeffotoni/goapiwebserver/goapiserver/config"
	"github.com/jeffotoni/goapiwebserver/goapiserver/pkg/util"
)

//
func ShowHtml(w http.ResponseWriter, r *http.Request) {

	// convert string to int
	WD_LEVEL, _ := strconv.Atoi(config.WD_LEVEL)

	//
	var tpl bytes.Buffer

	//
	type TmplHtml struct {
		Name string
	}

	//
	thtml := &TmplHtml{
		Name: "Go ApiServer",
	}

	// parse file html
	t := template.Must(template.ParseFiles(TemplateGwd("index.html", WD_LEVEL)))
	if err := t.Execute(&tpl, thtml); err != nil {
		util.Print("Error Template! [" + err.Error() + "]")
		os.Exit(0)
	}
	t.ExecuteTemplate(w, "index.html", &thtml)
}

func TemplateGwd(NameTmpl string, level int) string {
	pathNewOrg := GetWdLocal(level)
	pathNewOrg = pathNewOrg + "/templates/" + NameTmpl
	if util.FileExist(pathNewOrg) {

		return pathNewOrg
	}
	return ""
}

func GetWdLocal(level int) string {
	pathNow, _ := os.Getwd()
	pathV := strings.Split(pathNow, "/")
	pathNew := pathV[:len(pathV)-level]
	pathNewOrg := strings.Join(pathNew, "/")
	return pathNewOrg
}
