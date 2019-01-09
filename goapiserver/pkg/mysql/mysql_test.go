// Go Api server
// @jeffotoni
// 2019-01-04

package mysql

import (
	"reflect"
	"sync"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_cache_put(t *testing.T) {
	type fields struct {
		mm    sync.Map
		Mutex sync.Mutex
	}
	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cache{
				mm:    tt.fields.mm,
				Mutex: tt.fields.Mutex,
			}
			c.put(tt.args.key, tt.args.value)
		})
	}
}

func Test_cache_get(t *testing.T) {
	type fields struct {
		mm    sync.Map
		Mutex sync.Mutex
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cache{
				mm:    tt.fields.mm,
				Mutex: tt.fields.Mutex,
			}
			if got := c.get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cache.get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cache_loadStore(t *testing.T) {
	type fields struct {
		mm    sync.Map
		Mutex sync.Mutex
	}
	type args struct {
		key interface{}
		fc  func() interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cache{
				mm:    tt.fields.mm,
				Mutex: tt.fields.Mutex,
			}
			if got := c.loadStore(tt.args.key, tt.args.fc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cache.loadStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnect(t *testing.T) {
	tests := []struct {
		name string
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Connect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
