// Back-End in Go server
// @jeffotoni
// 2019-01-04

package ruser

import (
	"database/sql"
	"encoding/json"

	my "github.com/jeffotoni/goapiwebserver/goapiserver/pkg/mysql"
)

// search user
func GetFind(email string) string {
	// SUCESS
	var Db = my.MyDb.Mydb
	// Db...
	if interf := my.Connect(); interf != nil {
		Db = interf.(*sql.DB)
	} else {
		return ""
	}

	// crypt password
	// var adlogin = mlogin.AdLoginAuth{}
	// Buscarno banco extension
	row := Db.QueryRow("SELECT logi_firstname, logi_lastname, logi_phone, logi_email FROM ad_login WHERE logi_email = ?", email)
	var logi_firstname, logi_lastname, logi_phone, logi_email string
	errqy := row.Scan(&logi_firstname, &logi_lastname, &logi_phone, &logi_email)
	if errqy != nil {
		return ""
	}
	// user valid
	if len(logi_email) > 0 {
		type dataUser struct {
			Firstname string
			Lastname  string
			Phone     string
			Email     string
		}

		var u = dataUser{
			Firstname: logi_firstname,
			Lastname:  logi_lastname,
			Phone:     logi_phone,
			Email:     logi_email,
		}

		// create json
		userJson, err := json.Marshal(u)
		if err != nil {
			return ""
		}

		return string(userJson)
	}

	return ""
}

// it's user exist
func ExistUser(email string) bool {
	// SUCESS
	var Db = my.MyDb.Mydb
	// Db...
	if interf := my.Connect(); interf != nil {
		Db = interf.(*sql.DB)
	} else {
		return false
	}

	// crypt password
	// var adlogin = mlogin.AdLoginAuth{}
	// Buscarno banco extension
	row := Db.QueryRow("SELECT count(logi_id) as count FROM ad_login WHERE logi_email = ?", email)
	var count int
	errqy := row.Scan(&count)
	if errqy != nil {
		return false
	}
	// user valid
	if count > 0 {
		return true
	}

	return false
}
