// Back-End in Go server
// @jeffotoni
// 2019-01-04

package assets

import (
	"html/template"
	"net/http"
	"testing"
)

func TestGoMustAssetString(t *testing.T) {
	type args struct {
		strtext string
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
			if got := GoMustAssetString(tt.args.strtext); got != tt.want {
				t.Errorf("GoMustAssetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRender(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		r    *http.Request
		tpl  *template.Template
		name string
		data interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Render(tt.args.w, tt.args.r, tt.args.tpl, tt.args.name, tt.args.data)
		})
	}
}

func TestPush(t *testing.T) {
	type args struct {
		w        http.ResponseWriter
		resource string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Push(tt.args.w, tt.args.resource)
		})
	}
}
