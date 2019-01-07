// Go Api server
// @jeffotoni
// 2019-01-04

package handler

type Endpoint struct {
	Ping      string
	PostToken string
	//PutDoc  string
	PostLogin string
}

func SetEndPoint() *Endpoint {

	point := &Endpoint{
		Ping:      "/api/v1/ping",  //
		PostToken: "/api/v1/token", //
		PostLogin: "/api/v1/login", //
	}
	return point
}
