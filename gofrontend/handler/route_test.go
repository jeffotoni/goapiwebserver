// Front-end in Go server
// @jeffotoni
// 2019-01-04

package handler

import (
	"net/http"
	"reflect"
	"sync"
	"testing"

	"github.com/jeffotoni/goapiwebserver/gofrontend/config"
)

func TestStart(t *testing.T) {
	type args struct {
		cfg config.Config
	}
	tests := []struct {
		name string
		args args
		want *FrontEndServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Start(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrontEndServer_Stop(t *testing.T) {
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
			FrontEndServer := &FrontEndServer{
				server: tt.fields.server,
				wg:     tt.fields.wg,
			}
			if err := FrontEndServer.Stop(); (err != nil) != tt.wantErr {
				t.Errorf("FrontEndServer.Stop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
