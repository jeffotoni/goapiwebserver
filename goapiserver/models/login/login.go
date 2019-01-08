// Go Api server
// @jeffotoni
// 2019-01-04

package mlogin

type AdLoginAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginJson struct {
}
