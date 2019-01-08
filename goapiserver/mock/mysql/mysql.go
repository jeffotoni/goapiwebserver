// Go Api server
// @jeffotoni
// 2019-01-04

package main

import (
	"database/sql"
	"log"

	my "github.com/jeffotoni/goapiwebserver/goapiserver/pkg/mysql"
)

// testing mysql connection,
// and select login table
func main() {

	// SUCESS
	var Db = my.MyDb.Mydb

	// Db...
	if interf := my.Connect(); interf != nil {
		Db = interf.(*sql.DB)
	} else {

		jsonstr := `{"status":"error","message":"error ao fazer connect Mysql com Db.."}`
		log.Println(jsonstr)
		return
	}

	idUser := 1

	row := Db.QueryRow("SELECT count(logi_id) as count FROM ad_login WHERE logi_id=?", idUser)
	var countx int
	errqy := row.Scan(&countx)
	if errqy != nil {

		jsonstr := `{"status":"error","message":"Error QueryRow ad_login"}`
		log.Println(jsonstr)
		return
	}

	if countx > 0 {
		jsonstr := `{"status":"success","message":"Hello, exist"}`
		log.Println(jsonstr)
		return
	}
}
