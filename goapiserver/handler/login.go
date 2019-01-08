package handler

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	logg "github.com/jeffotoni/goapiwebserver/goapiserver/logg"
)

func Login(w http.ResponseWriter, r *http.Request) {

	s1 := logg.Start()

	if strings.ToUpper(r.Method) == "POST" {

		r.ParseForm()
		fmt.Println(r.Form)

		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		// w.WriteHeader(http.StatusMethodNotAllowed)
		// jsonstr := `{"status":"error","message":"Error the method has to be POST or GET!"}`
		// io.WriteString(w, jsonstr)
		logg.Show(SetEndPoint().GetLogin, strings.ToUpper(r.Method), "success", s1)
		return

	} else {

		w.WriteHeader(http.StatusMethodNotAllowed)
		jsonstr := `{"status":"error","message":"Error the method has to be POST or GET!"}`
		io.WriteString(w, jsonstr)
		logg.Show(SetEndPoint().Ping, strings.ToUpper(r.Method), "error", s1)
		return
	}
}
