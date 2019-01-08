package repol

import (
	"database/sql"

	my "github.com/jeffotoni/goapiwebserver/goapiserver/pkg/mysql"
)

func ValidLogin(email, password string) bool {
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
	// Get user, Exist or Not
	row := Db.QueryRow("SELECT count(logi_id) as count FROM ad_login WHERE logi_email = ? AND logi_password = ?", email, password)
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
