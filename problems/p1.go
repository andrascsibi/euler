package main

import (
	"fmt"
	"os"
	"strconv"
)

func P1(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func P2(n int) int {
	sum := 0
	for a, b := 1, 2; b < n; a, b = b, a+b {
		if b%2 == 0 {
			sum += b
		}
	}
	return sum
}

type Problem struct {
	solution func(int) int
	n        int
}

var problems = map[int]Problem{
	1: {P1, 1000},
	2: {P2, 4e6},
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
