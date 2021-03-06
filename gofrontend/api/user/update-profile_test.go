// Front-end in Go server
// @jeffotoni
// 2019-01-07

package user

import "testing"

func TestUpdateProfile(t *testing.T) {
	type args struct {
		token     string
		firstname string
		lastname  string
		phone     string
		email     string
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
			if got := UpdateProfile(tt.args.token, tt.args.firstname, tt.args.lastname, tt.args.phone, tt.args.email); got != tt.want {
				t.Errorf("UpdateProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
