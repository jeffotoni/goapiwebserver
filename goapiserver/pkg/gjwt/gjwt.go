// Go Api server
// @jeffotoni
// 2019-01-04

package gjwt

import (
	"crypto/rsa"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jeffotoni/goapiwebserver/goapiserver/cert"
	"github.com/jeffotoni/goapiwebserver/goapiserver/models/jwt"

	"net/http"
	"strconv"
	"strings"
	"time"
)

var (

	// default, we will load
	// the keys into memory
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey

	// only if it will be used by files
	pathPrivate = "./private.rsa"
	pathPublic  = "./public.rsa.pub"

	ProjectTitle = "jwt Goapiserver"

	ExpirationHours = 24 // Hours
	DayExpiration   = 30 // Days

	// md5(userR) = 123456ajeffotoni2020
	//  => 16d25077bf7365532ba4fe539619b443
	UserR = "16d25077bf7365532ba4fe539619b443"

	// md5(PassR) = 123456#2020
	// => ff7ca77e69408bc3c3d7279e38364b4a
	PassR = "ff7ca77e69408bc3c3d7279e38364b4a"
)

// Structure of our server configurations
type JsonMsg struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

// jwt init
func init() {

	var errx error
	privateByte := []byte(cert.RSA_PRIVATE)
	privateKey, errx = jwt.ParseRSAPrivateKeyFromPEM(privateByte)
	if errx != nil {
		return
	}

	publicByte := []byte(cert.RSA_PUBLIC)
	publicKey, errx = jwt.ParseRSAPublicKeyFromPEM(publicByte)
	if errx != nil {
		return
	}
}

// jwt generateJwt
func generateJwt(model models.User) (string, string) {

	// convert int to duration
	HourxMonth := time.Duration(ExpirationHours * DayExpiration)

	// Generating date validation to return to the user
	Expires := time.Now().Add(time.Hour * HourxMonth).Unix()

	// convert int6
	ExpiresInt64, _ := strconv.ParseInt(fmt.Sprintf("%v", Expires), 10, 64)

	// convert time unix to Date RFC
	ExpiresDateAll := time.Unix(ExpiresInt64, 0)

	// Date
	ExpiresDate := ExpiresDateAll.Format("2009-09-02")

	// claims Token data, the header
	claims := models.Claim{
		User: model.Login,
		StandardClaims: jwt.StandardClaims{
			// Expires in 24 hours * 10 days
			ExpiresAt: Expires,
			Issuer:    ProjectTitle,
		},
	}

	// Generating token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Transforming into string
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "Could not sign the token!", "2006-01-02"
	}

	// return token string
	return tokenString, ExpiresDate
}

//
func AuthBasic(w http.ResponseWriter, r *http.Request) {

	// Authorization Basic base64 Encode
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		//
		HttpWriteJson(w, "error", http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	//
	tokenBase64 := strings.Trim(auth[1], " ")

	//
	tokenBase64 = strings.TrimSpace(tokenBase64)

	// token 64
	authToken64 := strings.SplitN(tokenBase64, ":", 2)
	if len(authToken64) != 2 || authToken64[0] == "" || authToken64[1] == "" {
		HttpWriteJson(w, "error", "token base 64 invalid!", http.StatusUnauthorized)
		return
	}

	//
	tokenUserEnc := authToken64[0]

	//
	keyUserEnc := authToken64[1]

	// User, Login byte
	tokenUserDecode, _ := b64.StdEncoding.DecodeString(tokenUserEnc)

	// key user byte
	keyUserDec, _ := b64.StdEncoding.DecodeString(keyUserEnc)

	// User, Login string
	tokenUserDecodeS := strings.TrimSpace(strings.Trim(string(tokenUserDecode), " "))

	// key user, string
	keyUserDecS := strings.TrimSpace(strings.Trim(string(keyUserDec), " "))

	// Validate user and password in the database
	if tokenUserDecodeS == UserR && keyUserDecS == PassR {
		var model models.User
		model.Login = tokenUserDecodeS
		//model.Password = keyUserDec
		model.Password = ""
		model.Role = "admin"

		token, expires := generateJwt(model)
		result := models.ResponseToken{token, expires}
		jsonResult, err := json.Marshal(result)

		if err != nil {
			HttpWriteJson(w, "error", "json.Marshal error generating!", http.StatusUnauthorized)
			return
		}

		//
		w.WriteHeader(http.StatusOK)

		//
		w.Header().Set("Content-Type", "application/json")

		//
		w.Write(jsonResult)

		/**
		{
		  "Token": "39a3099b45634f6eb511991fddde83752_v2",
		  "Expires": "2026-09-14"
		}
		*/

	} else {

		//
		stringErr := "Invalid User or Key! " + auth[0] + " - " + auth[1]

		//
		w.WriteHeader(http.StatusForbidden)

		//
		w.Header().Set("Content-Type", "application/json")

		//
		HttpWriteJson(w, "error", stringErr, http.StatusUnauthorized)
	}

	defer r.Body.Close()
}

// AuthJwt
func AuthJwt(w http.ResponseWriter, r *http.Request) bool {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Bearer" {
		HttpWriteJson(w, "error", http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return false
	}

	token := strings.Trim(auth[1], " ")
	strings.TrimSpace(token)

	// star
	parsedToken, err := jwt.ParseWithClaims(token, &models.Claim{}, func(*jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil || !parsedToken.Valid {
		HttpWriteJson(w, "error", "Your token has expired!", http.StatusAccepted)
		return false

	}

	claims, ok := parsedToken.Claims.(*models.Claim)
	if !ok || claims.User != UserR {
		msgjson := `{"status":"error","message":"There's something strange about your token"}`
		io.WriteString(w, msgjson)
		return false
	}

	return true
}

// TokenGlobal = token
// UserGlobal = claims.User
// ExpiGlobal = fmt.Sprintf("%d", claims.ExpiresAt)
// fmt.Println("User: ", claims.User)
// func2(w, r)
func GetSplitTokenJwt(w http.ResponseWriter, r *http.Request) string {
	var Authorization string
	Authorization = r.Header.Get("Authorization")
	if Authorization == "" {
		Authorization = r.Header.Get("authorization")
	}
	// browsers
	if Authorization == "" {
		Authorization = r.Header.Get("Access-Control-Allow-Origin")
	}

	if Authorization != "" {
		auth := strings.SplitN(Authorization, " ", 2)
		if len(auth) != 2 || strings.TrimSpace(strings.ToLower(auth[0])) != "bearer" {
			return ""
		}

		token := strings.Trim(auth[1], " ")
		token = strings.TrimSpace(token)
		return token
	} else {
		return ""
	}
}

func CheckJwt(w http.ResponseWriter, r *http.Request) bool {
	if !tokenJwtClaimsValid(w, r) {
		return false
	}
	return true
}

func tokenJwtClaimsValid(w http.ResponseWriter, r *http.Request) bool {
	token := GetSplitTokenJwt(w, r)
	if token != "" {
		// start
		parsedToken, err := jwt.ParseWithClaims(token, &models.Claim{}, func(*jwt.Token) (interface{}, error) {
			return publicKey, nil
		})

		if err != nil || !parsedToken.Valid {
			return false
		}

		_, ok := parsedToken.Claims.(*models.Claim)
		return ok
	} else {
		return false
	}
}
