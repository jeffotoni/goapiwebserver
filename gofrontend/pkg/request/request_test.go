// Front-end in Go server
// @jeffotoni
// 2019-01-04

package Req

import (
	"io"
	"net/http"
	"reflect"
	"testing"
)

// test method HttpRequest
func TestHttpRequest(t *testing.T) {
	type args struct {
		method       string
		url          string
		content_type string
		body         io.Reader
	}
	tests := []struct {
		name     string
		args     args
		wantResp *http.Response
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := HttpRequest(tt.args.method, tt.args.url, tt.args.content_type, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("HttpRequest() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
