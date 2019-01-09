// Front-end in Go server
// @jeffotoni
// 2019-01-07

package login

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	tok "github.com/jeffotoni/goapiwebserver/gofrontend/api/token"
	//tok "github.com/jeffotoni/goapiwebserver/gofrontend/api/token"
)

// Let's call the login endpoint
func LoginValid(token, email, password string) string {
	// Url and endpoint
	apiUrl := tok.API_HOST_SERVER

	// only endpoint
	resource := tok.ApiEndpoint().PostGetLogin

	// cancel if you pass the team
	ctx, cancel := context.WithCancel(context.TODO())
	afterFuncTimer := time.AfterFunc(3*time.Second, func() { cancel() })
	defer afterFuncTimer.Stop()

	data := url.Values{}
	data.Set("email", email)
	data.Set("password", password)

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	req = req.WithContext(ctx)

	req.Header.Set("Authorization", "Bearer "+tok.ObjectToken())
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return string("error")
	}
	defer resp.Body.Close()

	if resp.Status == "200 OK" {
		type MessageJson struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		var mJson = MessageJson{}
		bodyMsg, _ := ioutil.ReadAll(resp.Body)
		if err := json.Unmarshal(bodyMsg, &mJson); err != nil {
			return string("error")
		}
		// check if success
		if strings.ToLower(mJson.Status) == "success" {
			return string("success")
		} else {
			return string("error")
		}
	} else {

		// try again 2 times
		return string("error")
	}
}
