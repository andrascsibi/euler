package problems

import (
	"fmt"
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

func TestProblems(t *testing.T) {
	var tests = []struct {
		probId, want int
	}{
		{1, 233168},
		{2, 4613732},
		{3, 6857},
		{4, 906609},
		{5, 232792560},
		{6, 25164150},
	}
	covered := make(map[int]bool)
	for _, c := range tests {
		covered[c.probId] = true
		problem := problems[c.probId]
		fmt.Printf("Solving problem #%v\n", c.probId)
		if got := problem.Solve(); got != c.want {
			t.Errorf("failed problem %v. wanted %v but got %v", c.probId, c.want, got)
		}
	}
	for probId, _ := range problems {
		if !covered[probId] {
			t.Errorf("problem #%v is not covered by test", probId)
		}
	}
}
