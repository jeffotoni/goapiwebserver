// Go Api server
// @jeffotoni
// 2019-01-04

package handler

import (
	"reflect"
	"testing"
)

func TestSetEndPoint(t *testing.T) {
	tests := []struct {
		name string
		want *Endpoint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetEndPoint(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetEndPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
