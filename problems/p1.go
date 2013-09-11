package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type Problem struct {
	n        int
	solution func(int) int
}

func Palindrome(n int) bool {
	switch {
	case n < 0:
		return false
	case n == 0:
		return true
	}
	digits := make([]uint8, 0, 10)
	for ; n > 0; n /= 10 {
		// least significant digit comes first
		digits = append(digits, uint8(n%10))
	}
	l := len(digits)
	for i := 0; i < l/2; i++ {
		if digits[i] != digits[l-i-1] {
			return false
		}
	}
	return true
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
				fmt.Println(i)
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
				if prod > greatest && Palindrome(prod) {
					greatest = prod
				}
			}
		}
		return greatest
	}},
	5: {20, func(n int) int {
		// I'm getting the impression that Go is not really a functional language:)
		all := func(list []int, predicate func(int) bool) bool {
			for _, item := range list {
				if !predicate(item) {
					return false
				}
			}
			return true
		}
		seq := func(from, to int) []int {
			l := to - from
			if l < 0 {
				l = 0
			}
			ret := make([]int, l)
			for i := 0; i < l; i++ {
				ret[i] = i + from
			}
			return ret
		}
		for i := n; ; i += n {
			if all(seq(1, n+1), func(j int) bool {
				return i%j == 0
			}) {
				return i
			}
		}
	}},
}

func main() {
	problemNumber, _ := strconv.Atoi(os.Args[1])
	problem, found := problems[problemNumber]
	if !found {
		fmt.Println("Problem", problemNumber, "is not solved yet.")
	}
	if len(os.Args) > 2 {
		problem.n, _ = strconv.Atoi(os.Args[2])
	}
	fmt.Println("Solving problem", problemNumber, "where n ==", problem.n)
	fmt.Println(problem.solution(problem.n))
}
