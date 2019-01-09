package handler

import (
	"io"
	"net/http"
	"strings"

	logg "github.com/jeffotoni/goapiwebserver/goapiserver/logg"
	"github.com/jeffotoni/goapiwebserver/goapiserver/repo/login"
)

func User(w http.ResponseWriter, r *http.Request) {
	s1 := logg.Start()
	if strings.ToUpper(r.Method) == "POST" {
		r.ParseForm()

		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		phone := r.FormValue("phone")
		email := r.FormValue("email")
		password := r.FormValue("password")

		if repol.CretaeNew(firstname, lastname, phone, email, password) {
			// Do createUser
			w.WriteHeader(http.StatusOK)
			jsonstr := `{"status":"success","message":"create user new with success!"}`
			io.WriteString(w, jsonstr)
			logg.Show(SetEndPoint().PostGetUser, strings.ToUpper(r.Method), "success", s1)
			return
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			jsonstr := `{"status":"error","message":""}`
			io.WriteString(w, jsonstr)
			logg.Show(SetEndPoint().PostGetUser, strings.ToUpper(r.Method), "error", s1)
			return
		}

	} else {
		// redirect register
		w.WriteHeader(http.StatusMethodNotAllowed)
		jsonstr := `{"status":"error","message":"Error the method has to be POST!"}`
		io.WriteString(w, jsonstr)
		logg.Show(SetEndPoint().PostGetUser, strings.ToUpper(r.Method), "error", s1)
		return
	}
}
