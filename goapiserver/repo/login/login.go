// Back-End in Go server
// @jeffotoni
// 2019-01-04

package repol

import (
	"database/sql"

	"github.com/jeffotoni/goapiwebserver/goapiserver/pkg/crypt"
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
	row := Db.QueryRow("SELECT logi_password FROM ad_login WHERE logi_email = ?", email)
	var logi_password string
	errqy := row.Scan(&logi_password)
	if errqy != nil {
		return false
	}
	// user valid
	if len(logi_password) > 0 {
		if crypt.CheckBlowfish(password, logi_password) {
			return true
		} else {
			return false
		}
	}

	return false
}
