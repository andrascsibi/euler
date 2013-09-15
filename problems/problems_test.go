package problems

import (
	"fmt"
	"testing"
)

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
		{7, 104743},
		//		{8, 40824}, // XXX inputs doesn't work because of relative path
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
			t.Logf("Warning! problem #%v is not covered by test", probId)
		}
	}
}
