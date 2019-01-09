// Front-end in Go server
// @jeffotoni
// 2019-01-07

package token

import (
	"context"
	"encoding/json"
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

func ObjectToken() string {
	// {"token":"xxx","expires":"2019-02-06","message":"use the token to access the endpoints"}
	byt := []byte(GetToken())
	var jsonToken = TokenStruct{}
	if err := json.Unmarshal(byt, &jsonToken); err != nil {
		// clean token
		SetToken("")
	}
	return jsonToken.Token
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

func FindToken() string {

	jsonToken := GetTokenApiServer()

	// exist, use
	if jsonToken != "error" {

		//struct message token server
		var stToken = &TokenStruct{}
		json.Unmarshal([]byte(jsonToken), &stToken)

		// token exists, can continue
		if stToken.Token != "" && len(stToken.Expires) == 10 {
			return stToken.Token
		}
	}

	return ""
}
