// Front-end in Go server
// @jeffotoni
// 2019-01-04

package util

import "testing"

func TestFileExist(t *testing.T) {
	type args struct {
		name string
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
			if got := FileExist(tt.args.name); got != tt.want {
				t.Errorf("FileExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirExist(t *testing.T) {
	type args struct {
		path string
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
			if got := DirExist(tt.args.path); got != tt.want {
				t.Errorf("DirExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateDirIfNotExist(t *testing.T) {
	type args struct {
		dir string
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
			if got := CreateDirIfNotExist(tt.args.dir); got != tt.want {
				t.Errorf("CreateDirIfNotExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateDirIfNotExistNotFile(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateDirIfNotExistNotFile(tt.args.dir)
		})
	}
}

func TestRemoveAllDir(t *testing.T) {
	type args struct {
		pathlocal string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveAllDir(tt.args.pathlocal)
		})
	}
}

func TestDeleteFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteFile(tt.args.path)
		})
	}
}
