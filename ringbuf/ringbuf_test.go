package ringbuf

import (
	"github.com/andrascsibi/euler/fun"
	"testing"
)

func TestRingBuf(t *testing.T) {

	var tests = []struct {
		newElem int
		wantSum int
	}{
		{1, 1},
		{2, 3},
		{3, 6},
		{4, 9},
		{10, 17},
	}

	rb := New(3)

	for _, c := range tests {
		rb.Add(c.newElem)
		if got := rb.Reduce(fun.Sum, 0); got != c.wantSum {
			t.Errorf("added elem %v. wanted %v. got %v", c.newElem, c.wantSum, got)
		}
	}
}
