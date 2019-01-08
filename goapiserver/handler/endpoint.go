// Go Api server
// @jeffotoni
// 2019-01-04

package handler

type Endpoint struct {
	Ping              string
	PostToken         string
	PostGetLogin      string
	PostSignup        string
	PostUpProfile     string
	PostForgotPssword string
	PostGetUser       string
}

func SetEndPoint() *Endpoint {

	point := &Endpoint{
		Ping:              "/api/v1/ping",            // POST/GET
		PostToken:         "/api/v1/token",           // GET
		PostGetLogin:      "/api/v1/login",           // POST/GET
		PostSignup:        "/api/v1/signup",          // POST
		PostUpProfile:     "/api/v1/profile",         // POST
		PostForgotPssword: "/api/v1/forgot-password", // POST
		PostGetUser:       "/api/v1/user",            // POST/GET
	}
	return point
}
