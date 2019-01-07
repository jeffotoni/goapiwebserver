// Go Api server
// @jeffotoni
// 2019-01-04

package models

// Claim structure, where we will use
// to validate our token with jwt

//
// jwt
//
type Claim struct {

	//
	//
	//
	User string `json:"user"`

	//
	//
	//
	jwt.StandardClaims
}
