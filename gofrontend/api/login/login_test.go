// Front-end in Go server
// @jeffotoni
// 2019-01-07

package login

import "testing"

func TestLoginValid(t *testing.T) {
	type args struct {
		token    string
		email    string
		password string
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
			if got := LoginValid(tt.args.token, tt.args.email, tt.args.password); got != tt.want {
				t.Errorf("LoginValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
