// Front-end in Go server
// @jeffotoni
// 2019-01-07

package login

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	tok "github.com/jeffotoni/goapiwebserver/gofrontend/api/token"
	//tok "github.com/jeffotoni/goapiwebserver/gofrontend/api/token"
)

type AdLogin struct {
	Logi_email    string `json:"logi_email"`
	Logi_password string `json:"logi_password"`
}

// Let's call the login endpoint
func Uservalid(token, user, password string) string {

	var Dlogin = AdLogin{Logi_email: user, Logi_password: password}

	byteLogin, err := json.Marshal(Dlogin)
	if err != nil {
		return string("error")
	}

	ctx, cancel := context.WithCancel(context.TODO())
	afterFuncTimer := time.AfterFunc(3*time.Second, func() { cancel() })
	defer afterFuncTimer.Stop()

	req, err := http.NewRequest("POST", tok.API_HOST_SERVER+tok.ApiEndpoint().GetLogin, bytes.NewBuffer(byteLogin))
	req = req.WithContext(ctx)

	req.Header.Set("Authorization", tok.AUTHORIZATION_BASIC)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return string("error")
	}
	defer resp.Body.Close()

	if resp.Status == "200 OK" {
		bodyDataLogin, _ := ioutil.ReadAll(resp.Body)
		return string(bodyDataLogin)
	} else {
		return string("error")
	}
}
