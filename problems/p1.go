package main

import (
	"fmt"
	"os"
	"strconv"
)

type Problem struct {
	n        int
	solution func(int) int
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
