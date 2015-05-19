package bitmask

import "testing"

func TestBitMask(t *testing.T) {
	tests := []struct {
		pos     int
		data    bool
		recheck bool
	}{
		{0, true, true},
		{1, true, true},
		{2, true, true},
		{3, false, false},
		{4, true, true},
		{5, true, false},
		{6, true, true},
		{7, true, true},
		{3, true, true},
		{5, false, true},
		{8, false, true},
		{9, true, true},
		{15, true, false},
		{15, false, true},
		{16, true, false},
		{16, false, false},
		{16, false, true},
		{31, true, true},
		{30, true, true},
		{29, true, true},
		{24, true, true},
	}
	b := New(32)
	for n, test := range tests {
		b.Set(test.pos, test.data)
		if got := b.Get(test.pos); got != test.data {
			t.Errorf("test 1-%d: expecting %v, got %v", n+1, test.data, got)
		}
	}
	for n, test := range tests {
		if !test.recheck {
			continue
		}
		if got := b.Get(test.pos); got != test.data {
			t.Errorf("test 2-%d: expecting %v, got %v", n+1, test.data, got)
		}
	}
}

func TestSetIfNot(t *testing.T) {
	tests := []struct {
		pos  int
		data bool
		ret  bool
	}{
		{0, true, true},
		{1, true, true},
		{2, true, true},
		{3, false, false},
		{4, true, true},
		{5, true, true},
		{6, true, true},
		{7, true, true},
		{3, true, true},
		{5, false, true},
		{8, false, false},
		{9, true, true},
		{15, true, true},
		{15, false, true},
		{16, true, true},
		{16, false, true},
		{16, false, false},
		{31, true, true},
		{30, true, true},
		{29, true, true},
		{24, true, true},
	}
	b := New(32)
	for n, test := range tests {
		if got := b.SetIfNot(test.pos, test.data); got != test.ret {
			t.Errorf("test d: expecting %v, got %v", n+1, test.ret, got)
		}
	}
}
