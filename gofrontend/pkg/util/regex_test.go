// Front-end in Go server
// @jeffotoni
// 2019-01-08

package util

import "testing"

func TestValidEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidEmail(tt.args.email); got != tt.want {
				t.Errorf("ValidEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
