// Go Api server
// @jeffotoni
// 2019-01-04

package config

import (
	"os"
	"strconv"
	"time"
)

const (
	LayoutDateLog = "2006-01-02 15:04:05"
	LayoutDate    = "2006-01-02"
	LayoutHour    = "15:04:05"
	MaxClients    = 20000
	NewLimiter    = 100
)

var (
	HOST_MAXBYTE         = os.Getenv("HOST_MAXBYTE")
	HOST_MAXBYTE_DEFAULT = 1 << 26 // 64MB

	HOST_SERVER         = os.Getenv("HOST_SERVER")
	HOST_SERVER_DEFAULT = "0.0.0.0:"

	PORT_SERVER         = os.Getenv("PORT_SERVER")
	PORT_SERVER_DEFAULT = "5002"

	X_KEY         = os.Getenv("X_KEY")
	AUTHORIZATION = os.Getenv("AUTHORIZATION")

	X_KEY_DEFAULT        = "MTIzNDU2"                            //123456
	AUTORIZATION_DEFAULT = "Bearer MTIzNDU2YWplZmZvdG9uaTIwMjA=" // 123456ajeffotoni2020

	READTIMEOUT  = 5 * time.Second
	WRITETIMEOUT = 10 * time.Second

	WD_LEVEL         = os.Getenv("WD_LEVEL")
	WD_LEVEL_DEFAULT = 0 // Path: Getwd - dir absolute
)

// Config provides
// basic configuration
type Config struct {
	Host           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

func Setenv() {

	if HOST_MAXBYTE == "" {
		// base64 => 1234
		HOST_MAXBYTE = strconv.Itoa(HOST_MAXBYTE_DEFAULT)
		os.Setenv("HOST_MAXBYTE", HOST_MAXBYTE)
	}

	if X_KEY == "" {
		// base64 => 1234
		X_KEY = X_KEY_DEFAULT
		os.Setenv("X_KEY", X_KEY)
	}

	if AUTHORIZATION == "" {
		// base64 => 1234AUTHORIZATION
		AUTHORIZATION = AUTORIZATION_DEFAULT
		os.Setenv("AUTHORIZATION", AUTHORIZATION)
	}

	if PORT_SERVER == "" {
		PORT_SERVER = PORT_SERVER_DEFAULT
		os.Setenv("PORT_SERVER", PORT_SERVER)
	}

	if HOST_SERVER == "" {
		HOST_SERVER = HOST_SERVER_DEFAULT + PORT_SERVER
		os.Setenv("HOST_SERVER", HOST_SERVER)
	} else {
		HOST_SERVER = HOST_SERVER + ":" + PORT_SERVER
	}

	if WD_LEVEL == "" {
		WD_LEVEL = strconv.Itoa(WD_LEVEL_DEFAULT)
		os.Setenv("WD_LEVEL", WD_LEVEL)
	}
}