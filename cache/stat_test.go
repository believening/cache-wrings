package cache

import "testing"

func TestStat_add(t *testing.T) {
	type args struct {
		k string
		v []byte
	}
	tests := []struct {
		name string
		s    *Stat
		args args
	}{
		{"add1", &Stat{10, 100, 1000}, args{"add1", []byte{'a', 'b', 'c'}}},
		{"add2",&Stat{1,4,4},args{"add1", []byte{'a', 'b', 'c','d'}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.add(tt.args.k, tt.args.v)
		})
	}
}

func TestStat_del(t *testing.T) {
	type args struct {
		k string
		v []byte
	}
	tests := []struct {
		name string
		s    *Stat
		args args
	}{
		{"add1", &Stat{10, 100, 1000}, args{"add1", []byte{'a', 'b', 'c'}}},
		{"add2",&Stat{1,4,4},args{"add1", []byte{'a', 'b', 'c','d'}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.del(tt.args.k, tt.args.v)
		})
	}
}
