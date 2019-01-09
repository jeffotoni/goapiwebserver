// Back-End in Go server
// @jeffotoni
// 2019-01-04

package repol

import "testing"

func TestValidLogin(t *testing.T) {
	type args struct {
		email    string
		password string
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
			if got := ValidLogin(tt.args.email, tt.args.password); got != tt.want {
				t.Errorf("ValidLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}
