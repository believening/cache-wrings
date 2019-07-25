package cache

import (
	"reflect"
	"testing"
)

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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInMemCache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInMemCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
