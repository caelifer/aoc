package main

import (
	"slices"
	"testing"
)

func TestLargestJolts(t *testing.T) {
	for _, tc := range []struct {
		charges []byte
		n       int
		want    []byte
	}{
		{[]byte{1, 2}, 1, []byte{2}},
		{[]byte{3, 1, 2}, 2, []byte{3, 2}},
	} {
		if got := largestJolts(tc.charges, tc.n); !slices.Equal(got, tc.want) {
			t.Errorf("given %+v, wanted %v, but got: %+v", tc.charges, tc.want, got)
		}
	}
}
