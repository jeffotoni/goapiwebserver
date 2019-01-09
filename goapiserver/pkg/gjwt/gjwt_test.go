// Go Api server
// @jeffotoni
// 2019-01-04

package gjwt

import (
	"net/http"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jeffotoni/ukkobox/models"
)

func Test_generateJwt(t *testing.T) {
	type args struct {
		model models.User
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := generateJwt(tt.args.model)
			if got != tt.want {
				t.Errorf("generateJwt() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("generateJwt() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCheckBasic(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name        string
		args        args
		wantOk      bool
		wantMsgjson string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, gotMsgjson := CheckBasic(tt.args.w, tt.args.r)
			if gotOk != tt.wantOk {
				t.Errorf("CheckBasic() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
			if gotMsgjson != tt.wantMsgjson {
				t.Errorf("CheckBasic() gotMsgjson = %v, want %v", gotMsgjson, tt.wantMsgjson)
			}
		})
	}
}

func TestGtokenJwt(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
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
			if got := GtokenJwt(tt.args.w, tt.args.r); got != tt.want {
				t.Errorf("GtokenJwt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSplitTokenBasic(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
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
			if got := GetSplitTokenBasic(tt.args.w, tt.args.r); got != tt.want {
				t.Errorf("GetSplitTokenBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckJwt(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
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
			if got := CheckJwt(tt.args.w, tt.args.r); got != tt.want {
				t.Errorf("CheckJwt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokenJwtClaimsValid(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
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
			if got := tokenJwtClaimsValid(tt.args.w, tt.args.r); got != tt.want {
				t.Errorf("tokenJwtClaimsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
