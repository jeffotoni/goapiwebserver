// Front-end in Go server
// @jeffotoni
// 2019-01-08

package repol

import (
	"database/sql"

	"github.com/jeffotoni/goapiwebserver/goapiserver/pkg/crypt"
	my "github.com/jeffotoni/goapiwebserver/goapiserver/pkg/mysql"
)

func CretaeNew(firstname, lastname, phone, email, password string) bool {
	// SUCESS
	var Db = my.MyDb.Mydb
	// Db...
	if interf := my.Connect(); interf != nil {
		Db = interf.(*sql.DB)
	} else {
		return false
	}

	// crypt blowfish
	password = crypt.Blowfish(password)

	// insert
	sqlStatement := `
INSERT INTO ad_login (logi_firstname,lastname,phone,email,password)
VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := Db.Exec(sqlStatement, lastname, phone, email, password)

	if err != nil {
		return false
	}

	return true
}
