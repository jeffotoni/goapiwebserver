// Go Api server
// @jeffotoni
// 2019-01-04

package middleware

import (
	"log"
	"reflect"
	"testing"
)

func TestNotify(t *testing.T) {
	type args struct {
		logger *log.Logger
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
			if got := Notify(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Notify() = %v, want %v", got, tt.want)
			}
		})
	}
}
