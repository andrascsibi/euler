package fun

import (
	"testing"
)

func TestPalindrome(t *testing.T) {

	var tests = []struct {
		n    int
		want bool
	}{
		{0, true},
		{9, true},
		{99, true},
		{999, true},
		{909, true},
		{-1, false},
		{10, false},
		{101, true},
		{1010101, true},
		{10100101, true},
		{10100001, false},
	}
	for i, c := range tests {
		if got := Palindrome(c.n); got != c.want {
			t.Errorf("failed case %v. got %v", i, got)
		}
	}
}

func TestReduce(t *testing.T) {
	var tests = []struct {
		slice   []int
		init    int
		reducer Reducer
		want    int
	}{
		{[]int{}, 0, Sum, 0},
		{[]int{1}, 0, Sum, 1},
		{[]int{1, 2, 3, 4}, 0, Sum, 10},
		{[]int{1, 2, 3, 4}, 10, Sum, 20},
		{[]int{1, 2, 3, 4}, 1, Prod, 24},
		{[]int{1, 2, 3, 4}, 10, Prod, 240},
	}

	for i, c := range tests {
		if got := Reduce(c.slice, c.reducer, c.init); got != c.want {
			t.Errorf("failed case %v. got %v", i, got)
		}
	}
}

func TestMemoizer(t *testing.T) {
	var fact Mapper
	fact = func(n int) int {
		if n == 1 {
			return 1
		}
		return n * fact(n-1)
	}
	if got := fact(5); got != 120 {
		t.Errorf("failed case")
	}
	memo := make([]int, 5)
	//	memoizer := Memoizer{memo, fact}
	//	memoizer.Run(5)
	fact = Memoize(memo, fact)
	if got := fact(5); got != 120 {
		t.Errorf("failed case")
	}
	fact(12)
}
