// Front-end in Go server
// @jeffotoni
// 2019-01-04

package Req

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/jeffotoni/goapiwebserver/gofrontend/config"
)

// responsible for triggering requests to apiserver
func HttpRequest(method, url, content_type string, body io.Reader) (resp *http.Response, err error) {
	if content_type == "" {
		content_type = "application/json"
	} else if strings.ToLower(content_type) == "json" {
		content_type = "application/json"
	}

	ctx, cancel := context.WithCancel(context.TODO())
	afterFuncTimer := time.AfterFunc(4*time.Second, func() { cancel() })
	defer afterFuncTimer.Stop()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	req = req.WithContext(ctx)
	//req.Header.Add("authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("X-Key", config.X_KEY)
	req.Header.Set("Authorization", config.AUTHORIZATION)
	req.Header.Set("Content-Type", content_type)

	client := http.Client{}
	resp, err = client.Do(req)
	//defer resp.Body.Close()
	return
}
