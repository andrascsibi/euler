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
