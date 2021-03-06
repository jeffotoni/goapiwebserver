// Front-end in Go server
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
)

var (
	HOST_MAXBYTE         = os.Getenv("HOST_MAXBYTE")
	HOST_MAXBYTE_DEFAULT = 1 << 26 // 64MB

	HOST_SERVER         = os.Getenv("HOST_SERVER")
	HOST_SERVER_DEFAULT = "0.0.0.0:"

	PORT_SERVER         = os.Getenv("PORT_SERVER")
	PORT_SERVER_DEFAULT = "5001"

	X_KEY         = os.Getenv("X_KEY")
	AUTHORIZATION = os.Getenv("AUTHORIZATION")
)

// Config provides basic configuration
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
		X_KEY = "MTIzNA=="
		os.Setenv("X_KEY", X_KEY)
	}

	if AUTHORIZATION == "" {
		// base64 => 1234AUTHORIZATION
		AUTHORIZATION = "MTIzNEFVVEhPUklaQVRJT04="
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
}
