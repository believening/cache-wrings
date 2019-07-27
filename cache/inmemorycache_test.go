package cache

import (
	"reflect"
	"testing"
)

// not consider Stat change
func Test_inMemoryCache_Set(t *testing.T) {
	type args struct {
		k string
		v []byte
	}
	tests := []struct {
		name    string
		imc     *inMemoryCache
		args    args
		wantErr bool
	}{
		{
			"set1",
			&inMemoryCache{m: make(map[string][]byte), Stat: Stat{0, 0, 0}},
			args{"set1", []byte{'s', 'e', 't', '1'}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.imc.Set(tt.args.k, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("inMemoryCache.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_inMemoryCache_Get(t *testing.T) {
	type args struct {
		k string
	}

	tests := []struct {
		name    string
		imc     *inMemoryCache
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "GetNormal",
			imc: &inMemoryCache{
				m: func(k string, v []byte) map[string][]byte {
					b := make(map[string][]byte)
					b[k] = v
					return b
				}("GetNormal", []byte("OK")),
				Stat: Stat{1, 9, 2},
			},
			args:    args{"GetNormal"},
			want:    []byte("OK"),
			wantErr: false,
		},
		{
			name: "NotExist",
			imc: &inMemoryCache{
				m: func(k string, v []byte) map[string][]byte {
					b := make(map[string][]byte)
					b[k] = v
					return b
				}("testKey", []byte("OK")),
				Stat: Stat{1, 7, 2},
			},
			args:    args{"NotExist"},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.imc.Get(tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("inMemoryCache.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inMemoryCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

// not consider Stat change
func Test_inMemoryCache_Del(t *testing.T) {
	type args struct {
		k string
	}
	tests := []struct {
		name    string
		imc     *inMemoryCache
		args    args
		wantErr bool
	}{
		{
			name: "Delete",
			imc: &inMemoryCache{
				m: func(k string, v []byte) map[string][]byte {
					b := make(map[string][]byte)
					b[k] = v
					return b
				}("DelTest", []byte("OK")),
				Stat: Stat{1, 7, 2},
			},
			args:    args{"DetTest"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.imc.Del(tt.args.k); (err != nil) != tt.wantErr {
				t.Errorf("inMemoryCache.Del() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_inMemoryCache_GetStat(t *testing.T) {
	tests := []struct {
		name string
		imc  *inMemoryCache
		want Stat
	}{
		{
			name: "getStat",
			imc: &inMemoryCache{
				m: func(k string, v []byte) map[string][]byte {
					b := make(map[string][]byte)
					b[k] = v
					return b
				}("stattest", []byte("OK")),
				Stat: Stat{1, 8, 2},
			},
			want: Stat{1, 8, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.imc.GetStat(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inMemoryCache.GetStat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInMemCache(t *testing.T) {
	tests := []struct {
		name string
		want Cache
	}{
		{
			name: "TestNew",
			want: &inMemoryCache{
				m:    make(map[string][]byte),
				Stat: Stat{0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInMemCache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInMemCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
