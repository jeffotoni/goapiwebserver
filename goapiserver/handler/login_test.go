package handler

import (
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
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
			Login(tt.args.w, tt.args.r)
		})
	}
}

func TestSignIn(t *testing.T) {
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
			SignIn(tt.args.w, tt.args.r)
		})
	}
}

func TestLoginExist(t *testing.T) {
	type args struct {
		email string
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
			LoginExist(tt.args.email, tt.args.w, tt.args.r)
		})
	}
}
