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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.del(tt.args.k, tt.args.v)
		})
	}
}
