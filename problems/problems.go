package problems

import (
	"github.com/andrascsibi/euler/fun"

	"fmt"
	"math"
)

var _ = fmt.Println // silence unused import complain

type Problem struct {
	input  int
	solver func(int) int
}

func (prob *Problem) Solve() int {
	return prob.solver(prob.input)
}

func (prob *Problem) SetInput(input int) {
	prob.input = input
}

func (prob *Problem) Input() int {
	return prob.input
}

func Get(problemId int) (problem Problem, solved bool) {
	problem, solved = problems[problemId]
	return
}

var problems = map[int]Problem{
	1: {1000, func(n int) int {
		sum := 0
		for i := 1; i < n; i++ {
			if i%3 == 0 || i%5 == 0 {
				sum += i
			}
		}
		return sum
	}},
	2: {4e6, func(n int) int {
		sum := 0
		for a, b := 1, 2; b < n; a, b = b, a+b {
			if b%2 == 0 {
				sum += b
			}
		}
		return sum
	}},
	3: {600851475143, func(n int) int {
		for i := 2; i < int(math.Sqrt(float64(n))); i++ {
			for ; n%i == 0; n = n / i {
			}
		}
		return n
	}},
	4: {3, func(n int) int {
		from := int(math.Pow10(n) - 1)
		to := int(math.Pow10(n - 1))
		greatest := 0
		for i := from; i > to; i-- {
			for j := from; j >= i; j-- {
				prod := i * j
				if prod > greatest && fun.Palindrome(prod) {
					greatest = prod
				}
			}
		}
		return greatest
	}},
	5: {20, func(n int) int {
		// I'm getting the impression that Go is not really a functional language:)
		for i := n; ; i += n {
			if fun.All(fun.Seq(1, n+1), func(j int) bool {
				return i%j == 0
			}) {
				return i
			}
		}
	}},
	6: {100, func(n int) int {
		sum := n * (n + 1) / 2
		sumOfSq := 0
		// or is it?
		for i := range fun.LazySeq(1, n+1) {
			sumOfSq += i * i
		}
		return sum*sum - sumOfSq
	}},
}
