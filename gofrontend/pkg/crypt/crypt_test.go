// Front-end in Go server
// @jeffotoni
// 2019-01-09

package crypt

import (
	"encoding/base64"
	"testing"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		textString string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.args.textString)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		textString string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.textString)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSeed(t *testing.T) {
	type args struct {
		n int
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
			if got := HashSeed(tt.args.n); got != tt.want {
				t.Errorf("HashSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandom(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Random(tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("Random() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandUid(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandUid(); got != tt.want {
				t.Errorf("RandUid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlowfish(t *testing.T) {
	type args struct {
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
			if got := Blowfish(tt.args.password); got != tt.want {
				t.Errorf("Blowfish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckBlowfish(t *testing.T) {
	type args struct {
		password string
		hash     string
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
			if got := CheckBlowfish(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("CheckBlowfish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGHash256(t *testing.T) {
	type args struct {
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
			if got := GHash256(tt.args.password); got != tt.want {
				t.Errorf("GHash256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGHash512(t *testing.T) {
	type args struct {
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
			if got := GHash512(tt.args.password); got != tt.want {
				t.Errorf("GHash512() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGSha1(t *testing.T) {
	type args struct {
		key string
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
			if got := GSha1(tt.args.key); got != tt.want {
				t.Errorf("GSha1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMd5(t *testing.T) {
	type args struct {
		text string
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
			if got := Md5(tt.args.text); got != tt.want {
				t.Errorf("Md5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HashFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUkkBase64Encode(t *testing.T) {
	type args struct {
		textString string
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
			if got := UkkBase64Encode(tt.args.textString); got != tt.want {
				t.Errorf("UkkBase64Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUkkBase64Decode(t *testing.T) {
	type args struct {
		textString string
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
			if got := UkkBase64Decode(tt.args.textString); got != tt.want {
				t.Errorf("UkkBase64Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
