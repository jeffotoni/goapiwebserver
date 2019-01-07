// Go Api server
// @jeffotoni
// 2019-01-04

package middleware

import (
	"reflect"
	"testing"
)

func TestMaxClients(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want Adapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxClients(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxClients() = %v, want %v", got, tt.want)
			}
		})
	}
}
