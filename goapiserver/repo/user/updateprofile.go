// Back-End in Go server
// @jeffotoni
// 2019-01-08

package ruser

import (
	"database/sql"

	my "github.com/jeffotoni/goapiwebserver/goapiserver/pkg/mysql"
)

// create new
func UpdateProfile(firstname, lastname, phone, email string) bool {

	// SUCESS
	var Db = my.MyDb.Mydb
	// Db...
	if interf := my.Connect(); interf != nil {
		Db = interf.(*sql.DB)
	} else {
		return false
	}

	// update
	sqlStatement := `
UPDATE ad_login SET logi_firstname = ?,logi_lastname = ?, logi_phone = ? WHERE logi_email = ?`
	_, err := Db.Exec(sqlStatement, firstname, lastname, phone, email)

	if err != nil {
		return false
	}
	return true
}
