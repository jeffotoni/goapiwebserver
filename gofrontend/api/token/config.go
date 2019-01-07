// Front-end in Go server
// @jeffotoni
// 2019-01-07

package token

import "os"

type ApiEndpoints struct {
	Ping              string
	PostToken         string
	GetLogin          string
	PostSignup        string
	PostUpProfile     string
	PostForgotPssword string
}

type TokenStruct struct {
	Token   string `json:"token"`
	Expires string `json:expires`
	Message string `json:Message`
}

var (
	API_SCHEME         = os.Getenv("API_SCHEME")
	API_SCHEME_DEFAULT = "http://"
	// SERVER
	API_URL         = os.Getenv("API_URL")
	API_URL_DEFAULT = "localhost"

	API_PORTA         = os.Getenv("API_PORTA")
	API_PORTA_DEFAULT = "5002"

	API_X_KEY         = os.Getenv("API_X_KEY")
	API_X_KEY_DEFAULT = "MTIzNDU2"

	AUTHORIZATION_SERVER  = os.Getenv("AUTHORIZATION_SERVER")
	AUTHORIZATION_DEFAULT = "MTIzNDU2YWplZmZvdG9uaTIwMjA="

	AUTHORIZATION_BASIC = ""

	API_HOST_SERVER = ""
)

func SetEnvKeys() {

	if API_X_KEY == "" {
		API_X_KEY = API_X_KEY_DEFAULT
		os.Setenv("API_X_KEY", API_X_KEY)
	}

	if AUTHORIZATION_SERVER == "" {
		AUTHORIZATION_SERVER = AUTHORIZATION_DEFAULT
		os.Setenv("AUTHORIZATION_SERVER", AUTHORIZATION_SERVER)
	}

	if API_SCHEME == "" {
		API_SCHEME = API_SCHEME_DEFAULT
		os.Setenv("API_SCHEME", API_SCHEME)
	}

	if API_URL == "" {
		API_URL = API_URL_DEFAULT
		os.Setenv("API_URL", API_URL)
	}

	if API_PORTA == "" {
		API_PORTA = API_PORTA_DEFAULT
		os.Setenv("API_PORTA", API_PORTA)
	}

	// create token access
	AUTHORIZATION_BASIC = "Basic " + API_X_KEY + ":" + AUTHORIZATION_SERVER
	API_HOST_SERVER = API_SCHEME + API_URL + ":" + API_PORTA
}

func ApiEndpoint() *ApiEndpoints {

	point := &ApiEndpoints{
		Ping:              "/api/v1/ping",            //
		PostToken:         "/api/v1/token",           //
		GetLogin:          "/api/v1/login",           //
		PostSignup:        "/api/v1/signup",          //
		PostUpProfile:     "/api/v1/profile",         //
		PostForgotPssword: "/api/v1/forgot-password", //
	}
	return point
}
