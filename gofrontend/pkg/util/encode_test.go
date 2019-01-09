// Front-end in Go server
// @jeffotoni
// 2019-01-04

package util

import "testing"

func TestEncodeBase64(t *testing.T) {
	type args struct {
		str string
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
			if got := EncodeBase64(tt.args.str); got != tt.want {
				t.Errorf("EncodeBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeBase64(t *testing.T) {
	type args struct {
		str64 string
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
			if got := DecodeBase64(tt.args.str64); got != tt.want {
				t.Errorf("DecodeBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}
