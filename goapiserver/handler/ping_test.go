// Go Api server
// @jeffotoni
// 2019-01-04

package handler

import (
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Ping(tt.args.w, tt.args.r)
		})
	}
}

func TestPingV2(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PingV2(tt.args.w, tt.args.r)
		})
	}
}
