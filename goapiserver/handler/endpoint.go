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
		Ping:      "/api/v1/ping",  // test server
		PostToken: "/api/v1/token", // insert
		PostLogin: "/api/v1/login", // select
	}
	return point
}
