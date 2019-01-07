// Go Api server
// @jeffotoni
// 2019-01-04

package handler

type Endpoint struct {
	Ping      string
	PostToken string
	//PutDoc  string
	GetLogin          string
	PostSignup        string
	PostUpProfile     string
	PostForgotPssword string
}

func SetEndPoint() *Endpoint {

	point := &Endpoint{
		Ping:              "/api/v1/ping",            //
		PostToken:         "/api/v1/token",           //
		GetLogin:          "/api/v1/login",           //
		PostSignup:        "/api/v1/signup",          //
		PostUpProfile:     "/api/v1/profile",         //
		PostForgotPssword: "/api/v1/forgot-password", //
	}
	return point
}
