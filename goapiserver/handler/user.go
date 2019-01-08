package handler

import (
	"io"
	"net/http"
	"strings"

	logg "github.com/jeffotoni/goapiwebserver/goapiserver/logg"
)

func User(w http.ResponseWriter, r *http.Request) {
	s1 := logg.Start()
	if strings.ToUpper(r.Method) == "POST" {
		// Do createUser
		// redirect register
		w.WriteHeader(http.StatusMethodNotAllowed)
		jsonstr := `{"status":"success","message":"create user new here."}`
		io.WriteString(w, jsonstr)
		logg.Show(SetEndPoint().PostGetLogin, strings.ToUpper(r.Method), "success", s1)
		return

	} else {
		// redirect register
		w.WriteHeader(http.StatusMethodNotAllowed)
		jsonstr := `{"status":"error","message":"Error the method has to be POST!"}`
		io.WriteString(w, jsonstr)
		logg.Show(SetEndPoint().PostGetLogin, strings.ToUpper(r.Method), "error", s1)
		return
	}
}
