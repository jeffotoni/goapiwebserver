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
		Ping:      "/v1/api/ping",  // test server
		PostToken: "/v1/api/token", // insert
		PostLogin: "/v1/api/login", // select
	}
	return point
}
