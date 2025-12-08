package main

import "testing"

func TestMatch(t *testing.T) {
	for _, tc := range []struct {
		tgt, pat string
		n int
		want     bool
	}{
		{"11", "1", 2, true},
		{"111", "1", 3, true},
		{"12341234", "1234", 2, true},
		{"123412345", "1234", 2, false},
		{"123451234", "1234", 2, false},
		{"23308757", "2330", 2, false},
	} {
		if got := match(tc.tgt, tc.pat, tc.n); got != tc.want {
			t.Errorf("for %s and %s: wanted %v, got %v", tc.tgt, tc.pat, tc.want, got)
		}
	}
}

func BenchmarkMatch(b *testing.B) {
	for b.Loop() {
		match("12341234", "1234", 2)
	}
}
