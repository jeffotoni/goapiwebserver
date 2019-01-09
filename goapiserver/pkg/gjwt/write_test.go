// Go Api server
// @jeffotoni
// 2019-01-04

package gjwt

import (
	"net/http"
	"testing"
)

func TestWriteJson(t *testing.T) {
	type args struct {
		Status string
		Msg    string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteJson(tt.args.Status, tt.args.Msg)
		})
	}
}

func TestHttpWriteJson(t *testing.T) {
	type args struct {
		w          http.ResponseWriter
		Status     string
		Msg        string
		httpStatus int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HttpWriteJson(tt.args.w, tt.args.Status, tt.args.Msg, tt.args.httpStatus)
		})
	}
}

func TestGetJson(t *testing.T) {
	type args struct {
		w          http.ResponseWriter
		Status     string
		Msg        string
		httpStatus int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetJson(tt.args.w, tt.args.Status, tt.args.Msg, tt.args.httpStatus); got != tt.want {
				t.Errorf("GetJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
