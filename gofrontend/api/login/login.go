// Front-end in Go server
// @jeffotoni
// 2019-01-07

package login

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

	//var Dlogin = AdLogin{Logi_email: user, Logi_password: password}
	// byteLogin, err := json.Marshal(Dlogin)
	// if err != nil {
	// 	return string("error")
	// }

	// Url and endpoint
	apiUrl := tok.API_HOST_SERVER

	// only endpoint
	resource := tok.ApiEndpoint().GetLogin

	// cancel if you pass the team
	ctx, cancel := context.WithCancel(context.TODO())
	afterFuncTimer := time.AfterFunc(3*time.Second, func() { cancel() })
	defer afterFuncTimer.Stop()

	data := url.Values{}
	data.Set("username", user)
	data.Set("password", password)

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	fmt.Println(urlStr)

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))
	req = req.WithContext(ctx)

	req.Header.Set("Authorization", tok.AUTHORIZATION_BASIC)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

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
