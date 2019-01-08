// Go Api server
// @jeffotoni
// 2019-01-04

package gjwt

import "crypto/rsa"

var (

	// default, we will load
	// the keys into memory
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey

	// only if it will be used by files
	PathPrivate = "./private.rsa"
	PathPublic  = "./public.rsa.pub"

	ProjectTitle = "jwt Goapiserver"

	ExpirationHours = 24 // Hours
	DayExpiration   = 30 // Days

	// base64 = MTIzNDU2
	// md5 =
	UserR = "123456"

	// base64 = MTIzNDU2YWplZmZvdG9uaTIwMjA=
	// md5 = 16d25077bf7365532ba4fe539619b443
	PassR = "123456ajeffotoni2020"
)

// Structure of our server configurations
type JsonMsg struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
