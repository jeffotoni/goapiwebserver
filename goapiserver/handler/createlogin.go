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

// content-type multipart/form-data
func SaveLogin(w http.ResponseWriter, r *http.Request) {

	s1 := logg.Start()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	msg := `{"status":"success",message":"everything is OK"}`

	// var Db = my.MyDb.Mydb
	// // Db...
	// if interf := my.Connect(); interf != nil {
	// 	Db = interf.(*sql.DB)
	// } else {
	// 	w.WriteHeader(http.StatusUnprocessableEntity)
	// 	jsonstr := `{"status":"error","msg":"error ao fazer connect Mysql com Db.."}`
	// 	io.WriteString(w, jsonstr)
	// 	logg.Show(SetEndPoint().PostApdata, strings.ToUpper(r.Method), "error", s1)
	// 	return
	// }

	// if r.Header.Get("X-key") != conf.X_KEY {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	jsonstr := `{"status":"error","msg":"Error de chave Key!"}`
	// 	io.WriteString(w, jsonstr)
	// 	logg.Show(SetEndPoint().PostApdata, strings.ToUpper(r.Method), "error", s1)
	// 	return
	// }

	if strings.ToUpper(r.Method) == "POST" {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, msg)
		logg.Show("endpoint/save", strings.ToUpper(r.Method), "everything is OK", s1)
		return
	} else {
		msg = `{"status":"erro",message":"something wrong"}`
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, msg)
		logg.Show("endpoint/save", strings.ToUpper(r.Method), "something wrong", s1)
		return
	}
}
