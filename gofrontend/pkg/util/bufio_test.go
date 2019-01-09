// Front-end in Go server
// @jeffotoni
// 2019-01-04

package util

import "testing"

func TestPrint(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Print(tt.args.text)
		})
	}
}
