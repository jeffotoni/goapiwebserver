// Back-End in Go server
// @jeffotoni
// 2019-01-08

package ruser

import "testing"

func TestCretaeNew(t *testing.T) {
	type args struct {
		firstname string
		lastname  string
		phone     string
		email     string
		password  string
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
			if got := CretaeNew(tt.args.firstname, tt.args.lastname, tt.args.phone, tt.args.email, tt.args.password); got != tt.want {
				t.Errorf("CretaeNew() = %v, want %v", got, tt.want)
			}
		})
	}
}
