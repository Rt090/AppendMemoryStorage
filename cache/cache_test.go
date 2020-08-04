package cache

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache_Insert(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		args []args
		want []string
	}{
		{
			name: "Insert when empty",
			args: []args{args{
				key:   "key1",
				value: "val1",
			}},
			want: []string{"val1"},
		},
		{
			name: "Insert when non empty - expect append",
			args: []args{args{
				key:   "key1",
				value: "val1",
			},
				{
					key:   "key1",
					value: "val2",
				}},
			want: []string{"val1", "val2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCache()
			for _, arg := range tt.args {
				c.Insert(arg.key, arg.value)
			}
			require.Equal(t, tt.want, c.Get(tt.args[0].key))
		})
	}
}

// func TestCache_Get(t *testing.T) {
// 	type args struct {
// 		key string
// 	}
// 	tests := []struct {
// 		name string
// 		c    *Cache
// 		args args
// 		want []string{}
// 	}{}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c :=
// 			if got := c.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Cache.Get() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestCache_Stats(t *testing.T) {
	type fields struct {
		buckets    []*bucket
		RWMutex    sync.RWMutex
		bucketSize uint32
		entries    uint32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				buckets:    tt.fields.buckets,
				RWMutex:    tt.fields.RWMutex,
				bucketSize: tt.fields.bucketSize,
				entries:    tt.fields.entries,
			}
			c.Stats()
		})
	}
}
