// Go Api server
// @jeffotoni
// 2019-01-04

package middleware

import (
	"reflect"
	"testing"
)

func TestGtokenJwt(t *testing.T) {
	tests := []struct {
		name string
		want Adapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GtokenJwt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GtokenJwt() = %v, want %v", got, tt.want)
			}
		})
	}
}
