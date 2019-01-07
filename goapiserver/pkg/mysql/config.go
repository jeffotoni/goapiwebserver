// Go Api server
// @jeffotoni
// 2019-01-04

package mysql

import "os"

var (
	DB_HOST     = os.Getenv("DB_HOST")
	DB_NAME     = os.Getenv("DB_NAME")
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_PORT     = os.Getenv("DB_PORT")
	DB_SSL      = ""
	DB_SORCE    = "mysql"
)
