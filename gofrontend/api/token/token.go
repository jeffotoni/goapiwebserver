// Front-end in Go server
// @jeffotoni
// 2019-01-07

package token

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

// token methods
// and variables
var (
	TOKEN_VALID = ""
)

func SetToken(token string) {
	TOKEN_VALID = token
}

func GetToken() string {
	return TOKEN_VALID
}

// client can use the token and access all available
// endpoints with more flexibii...
func GetTokenApiServer() string {

	// is there a token?
	// can search the server
	if tok := GetToken(); tok != "" {
		return GetToken()
	}

	// setting the keys
	// to generate the token
	SetEnvKeys()

	ctx, cancel := context.WithCancel(context.TODO())
	afterFuncTimer := time.AfterFunc(3*time.Second, func() { cancel() })
	defer afterFuncTimer.Stop()
	req, err := http.NewRequest("GET", API_HOST_SERVER+ApiEndpoint().PostToken, nil)
	req = req.WithContext(ctx)

	req.Header.Set("Authorization", AUTHORIZATION_BASIC)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {

		SetToken("")
		return string("error")
	}
	defer resp.Body.Close()

	//log.Println("response Status:", resp.Status)
	//log.Println("response Headers:", resp.Header)

	if resp.Status == "200 OK" {
		bodyToken, _ := ioutil.ReadAll(resp.Body)
		SetToken(string(bodyToken))
		return GetToken()

	} else {

		SetToken("")
		return string("error")
	}
}
