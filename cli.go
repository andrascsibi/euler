package main

import (
	"github.com/andrascsibi/euler/problems"

	"fmt"
	"os"
	"strconv"
)

type Problem struct {
	input  int
	solver func(int) int
}

func main() {
	problemId, _ := strconv.Atoi(os.Args[1])
	problem, solved := problems.Get(problemId)
	if !solved {
		fmt.Println("Problem", problemId, "is not solved yet.")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		input, _ := strconv.Atoi(os.Args[2])
		problem.SetInput(input)
	}
	fmt.Printf("Solving problem #%v where input == %v\n", problemId, problem.Input())
	fmt.Println(problem.Solve())
	os.Exit(0)
}
