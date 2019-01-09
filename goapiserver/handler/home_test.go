package handler

import (
	"net/http"
	"testing"
)

func Test_homeHandler(t *testing.T) {
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
			homeHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestEndpointRegex(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := EndpointRegex(tt.args.r)
			if got != tt.want {
				t.Errorf("EndpointRegex() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("EndpointRegex() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRegexEmail(t *testing.T) {
	type args struct {
		nameregex string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegexEmail(tt.args.nameregex); got != tt.want {
				t.Errorf("RegexEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
