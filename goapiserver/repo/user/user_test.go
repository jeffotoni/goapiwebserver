// Back-End in Go server
// @jeffotoni
// 2019-01-04

package ruser

import "testing"

func TestGetFind(t *testing.T) {
	type args struct {
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
			if got := GetFind(tt.args.email); got != tt.want {
				t.Errorf("GetFind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExistUser(t *testing.T) {
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
			if got := ExistUser(tt.args.email); got != tt.want {
				t.Errorf("ExistUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
