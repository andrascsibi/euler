package problems

import (
	"testing"
)

func TestP1(t *testing.T) {

	var tests = []struct {
		n, want int
	}{
		{10, 23},
		{1000, 233168},
	}
	for i, c := range tests {
		if got := P1(c.n); got != c.want {
			t.Errorf("failed case %v. got %v", i, got)
		}
	}
}
