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
	fmt.Println("Solving problem", problemNumber, "where n :=", problem.n)
	fmt.Println(problem.solution(problem.n))
}
