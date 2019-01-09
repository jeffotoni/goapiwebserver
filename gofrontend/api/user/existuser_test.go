// Front-end in Go server
// @jeffotoni
// 2019-01-07

package user

import "testing"

func TestExistUser(t *testing.T) {
	type args struct {
		token string
		email string
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
			if got := ExistUser(tt.args.token, tt.args.email); got != tt.want {
				t.Errorf("ExistUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
