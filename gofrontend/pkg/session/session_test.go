// Front-end in Go server
// @jeffotoni
// 2019-01-04

package session

import (
	"net/http"
	"testing"
)

func TestNameSession(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NameSession(); got != tt.want {
				t.Errorf("NameSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet(t *testing.T) {
	type args struct {
		session_name  string
		session_key   string
		session_value string
		w             http.ResponseWriter
		r             *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Set(tt.args.session_name, tt.args.session_key, tt.args.session_value, tt.args.w, tt.args.r)
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		session_name string
		session_key  string
		w            http.ResponseWriter
		r            *http.Request
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
			if got := Get(tt.args.session_name, tt.args.session_key, tt.args.w, tt.args.r); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDestroy(t *testing.T) {
	type args struct {
		session_name string
		session_key  string
		w            http.ResponseWriter
		r            *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Destroy(tt.args.session_name, tt.args.session_key, tt.args.w, tt.args.r)
		})
	}
}
