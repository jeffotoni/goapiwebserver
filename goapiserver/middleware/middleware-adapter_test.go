// Go Api server
// @jeffotoni
// 2019-01-04

package middleware

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAdapt(t *testing.T) {
	type args struct {
		h        http.Handler
		adapters []Adapter
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Adapt(tt.args.h, tt.args.adapters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Adapt() = %v, want %v", got, tt.want)
			}
		})
	}
}
