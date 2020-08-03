package cache

import (
	"reflect"
	"testing"
)

func TestCache_Insert(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		c    *Cache
		args args
	}{
		{
			name: "Test run",
			c:    &Cache{},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{}
			panic("oops")
			c.Insert(tt.args.key, tt.args.value)
		})
	}
}

func TestCache_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		c    *Cache
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{}
			if got := c.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
