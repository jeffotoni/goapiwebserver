// Go Api server
// @jeffotoni
// 2019-01-04

package handler

import (
	"io"
	"net/http"
	"strings"

	logg "github.com/jeffotoni/goapiwebserver/goapiserver/logg"
)

// check if the service is online
func Ping(w http.ResponseWriter, r *http.Request) {

	s1 := logg.Start()

	if strings.ToUpper(r.Method) != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		jsonstr := `{"status":"error","msg":"Error O método tem que ser POST!"}`
		io.WriteString(w, jsonstr)
		logg.Show(SetEndPoint().Ping, strings.ToUpper(r.Method), "error", s1)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	jsonstr := `{"status":"success","msg":"pong","method";"` + strings.ToUpper(r.Method) + `"}`
	io.WriteString(w, jsonstr)
	logg.Show(SetEndPoint().Ping, strings.ToUpper(r.Method), "success", s1)
	return
}

func PingV2(w http.ResponseWriter, r *http.Request) {

	s1 := logg.Start()

	if strings.ToUpper(r.Method) != "POST" && strings.ToUpper(r.Method) != "GET" && strings.ToUpper(r.Method) != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		jsonstr := `{"status":"error","msg":"Error The method has to be POST/GET OU PUT!"}`
		io.WriteString(w, jsonstr)
		logg.Show(SetEndPoint().Ping, strings.ToUpper(r.Method), "error", s1)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	jsonstr := `{"status":"ok","msg":"pong","method";"` + strings.ToUpper(r.Method) + `"}`
	io.WriteString(w, jsonstr)
	logg.Show(SetEndPoint().Ping, strings.ToUpper(r.Method), "sucesso", s1)
	return
}
