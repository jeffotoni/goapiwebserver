// Front-end in Go server
// @jeffotoni
// 2019-01-04

package handler

import (
	"net/http"
	"testing"
)

func TestAdminHandler(t *testing.T) {
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
			AdminHandler(tt.args.w, tt.args.r)
		})
	}
}
