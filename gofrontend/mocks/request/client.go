// Front-end in Go server
// @jeffotoni
// 2019-01-04

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/jeffotoni/goapiwebserver/gofrontend/pkg/request"
)

type ipTest struct{}

func main() {

	b, _ := json.Marshal(ipTest{})
	resp, err := post(http.MethodPost, "https://api.ipify.org?format=json", "json", bytes.NewReader(b))
	if err != nil {
		fmt.Println("Error: post")
		return
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ReadAll")
		return
	}

	fmt.Println("---------- shooting post ---------- ")
	fmt.Println("Data:", string(b))
	fmt.Println("status Client:", resp.Status)

	// get
	resp, err = get(http.MethodGet, "https://api.ipify.org?format=json", "json", nil)
	if err != nil {
		fmt.Println("Error: post")
		return
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ReadAll")
		return
	}

	fmt.Println("---------- shooting get ---------- ")
	fmt.Println("Data:", string(b))
	fmt.Println("status Client:", resp.Status)
}

func post(method, url, content_type string, body io.Reader) (resp *http.Response, err error) {

	resp, err = Req.HttpRequest(
		method,
		url,
		content_type,
		body,
	)

	return
}

func get(method, url, content_type string, body io.Reader) (resp *http.Response, err error) {

	return post(method, url, content_type, nil)
}
