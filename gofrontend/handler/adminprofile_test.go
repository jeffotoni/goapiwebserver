// Front-end in Go server
// @jeffotoni
// 2019-01-09

package handler

import (
	"net/http"
	"testing"
)

func TestAdminProfileHandler(t *testing.T) {
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
			AdminProfileHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_tplUserProfileHtml(t *testing.T) {
	type args struct {
		ruser *UserProfile
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
			tplUserProfileHtml(tt.args.ruser, tt.args.w, tt.args.r)
		})
	}
}
