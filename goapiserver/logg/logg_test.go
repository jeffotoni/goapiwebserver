// Go Api server
// @jeffotoni
// 2019-01-04

package logg

import (
	"reflect"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	tests := []struct {
		name   string
		wantS1 time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS1 := Start(); !reflect.DeepEqual(gotS1, tt.wantS1) {
				t.Errorf("Start() = %v, want %v", gotS1, tt.wantS1)
			}
		})
	}
}

func TestShow(t *testing.T) {
	type args struct {
		endpoint string
		metodo   string
		status   string
		s1       time.Time
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Show(tt.args.endpoint, tt.args.metodo, tt.args.status, tt.args.s1)
		})
	}
}
