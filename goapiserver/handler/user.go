package handler

import (
	"io"
	"net/http"
	"strings"

	logg "github.com/jeffotoni/goapiwebserver/goapiserver/logg"
	"github.com/jeffotoni/goapiwebserver/goapiserver/repo/user"
)

func User(w http.ResponseWriter, r *http.Request) {
	s1 := logg.Start()
	if strings.ToUpper(r.Method) == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusMethodNotAllowed)
			jsonstr := `{"status":"error","message":"parse error in form!"}`
			io.WriteString(w, jsonstr)
			logg.Show(SetEndPoint().PostGetUser, strings.ToUpper(r.Method), "error", s1)
			return
		}

		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		phone := r.FormValue("phone")
		email := r.FormValue("email")
		password := r.FormValue("password")

		if ruser.CretaeNew(firstname, lastname, phone, email, password) {
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

	} else if strings.ToUpper(r.Method) == http.MethodPut { // Update Profile
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusMethodNotAllowed)
			jsonstr := `{"status":"error","message":"parse error in form!"}`
			io.WriteString(w, jsonstr)
			logg.Show(SetEndPoint().PostGetUser, strings.ToUpper(r.Method), "error", s1)
			return
		}

		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		phone := r.FormValue("phone")
		email := r.FormValue("email")

		if ruser.UpdateProfile(firstname, lastname, phone, email) {
			// Do Update in Profile
			w.WriteHeader(http.StatusOK)
			jsonstr := `{"status":"success","message":"update user with success!"}`
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
		jsonstr := `{"status":"error","message":"Error the method has to be PUT!"}`
		io.WriteString(w, jsonstr)
		logg.Show(SetEndPoint().PostGetUser, strings.ToUpper(r.Method), "error", s1)
		return
	}
}

// get login/{email}
func FindGetUser(email string, w http.ResponseWriter, r *http.Request) {
	s1 := logg.Start()
	// valid email
	email = strings.TrimSpace(strings.ToLower(email))

	if len(email) == 0 {
		w.WriteHeader(http.StatusOK)
		jsonstr := `{"status":"error","message":"email can not be empty!"}`
		io.WriteString(w, jsonstr)
		logg.Show(SetEndPoint().PostGetUser, strings.ToUpper(r.Method), "error", s1)
		return
	}

	// valid email user
	if ruser.ExistUser(email) {
		// search user data
		jsonUser := ruser.GetFind(email)
		// convert json
		if jsonUser != "" {
			w.WriteHeader(http.StatusOK)
			//jsonstr := `{"status":"success","message":"all data in json"}`
			io.WriteString(w, jsonUser)
			logg.Show(SetEndPoint().PostGetUser, strings.ToUpper(r.Method), "success", s1)
			return
		}
	} else {
		w.WriteHeader(http.StatusOK)
		jsonstr := `{"status":"error","message":"user not exist"}`
		io.WriteString(w, jsonstr)
		logg.Show(SetEndPoint().PostGetUser, strings.ToUpper(r.Method), "error", s1)
		return
	}
}
