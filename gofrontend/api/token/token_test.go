// Front-end in Go server
// @jeffotoni
// 2019-01-07

package token

import "testing"

func TestSetToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetToken(tt.args.token)
		})
	}
}

func TestGetToken(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetToken(); got != tt.want {
				t.Errorf("GetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectToken(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ObjectToken(); got != tt.want {
				t.Errorf("ObjectToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTokenApiServer(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTokenApiServer(); got != tt.want {
				t.Errorf("GetTokenApiServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindToken(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindToken(); got != tt.want {
				t.Errorf("FindToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
