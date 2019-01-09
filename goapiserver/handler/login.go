package handler

import (
	"io"
	"net/http"
	"strings"

	logg "github.com/jeffotoni/goapiwebserver/goapiserver/logg"
	repol "github.com/jeffotoni/goapiwebserver/goapiserver/repo/login"
	"github.com/jeffotoni/goapiwebserver/goapiserver/repo/user"
	//repologin "github.com/jeffotoni/goapiwebserver/goapiserver/repo/login"
)

func Login(w http.ResponseWriter, r *http.Request) {
	s1 := logg.Start()
	if strings.ToUpper(r.Method) == "POST" {
		// Do login
		SignIn(w, r)
	} else {
		// redirect register
		w.WriteHeader(http.StatusMethodNotAllowed)
		jsonstr := `{"status":"error","message":"Error the method has to be POST or GET!"}`
		io.WriteString(w, jsonstr)
		logg.Show(SetEndPoint().PostGetLogin, strings.ToUpper(r.Method), "error", s1)
		return
	}
}

// Do login
func SignIn(w http.ResponseWriter, r *http.Request) {
	s1 := logg.Start()
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")
	// valid email
	email = strings.TrimSpace(strings.ToLower(email))

	if len(email) == 0 {
		w.WriteHeader(http.StatusOK)
		jsonstr := `{"status":"error","message":"Error email is is required"}`
		io.WriteString(w, jsonstr)
		logg.Show(SetEndPoint().PostGetLogin, strings.ToUpper(r.Method), "error", s1)
		return
	}

	if len(password) == 0 {
		w.WriteHeader(http.StatusOK)
		jsonstr := `{"status":"error","message":"Error Password is is required"}`
		io.WriteString(w, jsonstr)
		logg.Show(SetEndPoint().PostGetLogin, strings.ToUpper(r.Method), "error", s1)
		return
	}

	// authenticates the user in the bank
	if repol.ValidLogin(email, password) {
		w.WriteHeader(http.StatusOK)
		jsonstr := `{"status":"success","message":"found user"}`
		io.WriteString(w, jsonstr)
		logg.Show(SetEndPoint().PostGetLogin, strings.ToUpper(r.Method), "error", s1)
		return
	}

	jsonstr := `{"status":"error","message":"Error user could not authenticate"}`
	io.WriteString(w, jsonstr)
	logg.Show(SetEndPoint().PostGetLogin, strings.ToUpper(r.Method), "error", s1)
	return
}

// get login/{email}
func LoginExist(email string, w http.ResponseWriter, r *http.Request) {
	s1 := logg.Start()
	// valid email
	email = strings.TrimSpace(strings.ToLower(email))

	if len(email) == 0 {
		w.WriteHeader(http.StatusOK)
		jsonstr := `{"status":"error","message":"email can not be empty!"}`
		io.WriteString(w, jsonstr)
		logg.Show(SetEndPoint().PostGetLogin, strings.ToUpper(r.Method), "error", s1)
		return
	}

	if ruser.ExistUser(email) {
		w.WriteHeader(http.StatusOK)
		jsonstr := `{"status":"success","message":"user exist"}`
		io.WriteString(w, jsonstr)
		logg.Show(SetEndPoint().PostGetLogin, strings.ToUpper(r.Method), "success", s1)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		jsonstr := `{"status":"error","message":"user not exist"}`
		io.WriteString(w, jsonstr)
		logg.Show(SetEndPoint().PostGetLogin, strings.ToUpper(r.Method), "error", s1)
		return
	}
}
