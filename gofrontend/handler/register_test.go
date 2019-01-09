// Front-end in Go server
// @jeffotoni
// 2019-01-04

package handler

import (
	"net/http"
	"testing"
)

func TestLoginHandlerRegister(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoginHandlerRegister(tt.args.w, tt.args.r)
		})
	}
}

func Test_tplRegisterHtml(t *testing.T) {
	type args struct {
		ruser *RegisterUser
		w     http.ResponseWriter
		r     *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tplRegisterHtml(tt.args.ruser, tt.args.w, tt.args.r)
		})
	}
}
