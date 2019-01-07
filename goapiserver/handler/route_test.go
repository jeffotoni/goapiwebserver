// Go Api server
// @jeffotoni
// 2019-01-04

package handler

import (
	"net/http"
	"reflect"
	"sync"
	"testing"

	"github.com/jeffotoni/goapiwebserver/goapiserver/config"
)

func TestStartServer(t *testing.T) {
	type args struct {
		cfg config.Config
	}
	tests := []struct {
		name string
		args args
		want *GoServerHttp
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartServer(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoServerHttp_StopServer(t *testing.T) {
	type fields struct {
		server *http.Server
		wg     sync.WaitGroup
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GoServerHttp := &GoServerHttp{
				server: tt.fields.server,
				wg:     tt.fields.wg,
			}
			if err := GoServerHttp.StopServer(); (err != nil) != tt.wantErr {
				t.Errorf("GoServerHttp.StopServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSetenv(t *testing.T) {
	type args struct {
		cfg config.Config
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Setenv(tt.args.cfg)
		})
	}
}
