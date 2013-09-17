package problems

import (
	"github.com/andrascsibi/euler/fun"
	"github.com/andrascsibi/euler/inputs"

	"fmt"
	"math"
	"strconv"
	"bufio"
	"strings"
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
	7: {10001, func(n int) int {
		primes := make([]int, n)
		isPrime := func(k int) bool {
			to := int(math.Sqrt(float64(k)))
			for _, p := range primes {
				if p > to {
					return true
				}
				if k%p == 0 {
					return false
				}
			}
			return true
		}
		primes[0] = 2
		primes[1] = 3
		prevPrime := 3
		for i := 2; i < n; i++ {
			nextPrime := prevPrime + 2
			for ; !isPrime(nextPrime); nextPrime += 2 {
			}
			primes[i] = nextPrime
			prevPrime = nextPrime
		}
		return primes[n-1]
	}},
	8: {0, func(n int) int {
		br, closer := inputs.BufReader("inputs", 8, n)
		defer closer.Close()
		toDigit := func(b byte) int8 {
			if b < '0' || b > '9' {
				return -1
			}
			return int8(b - '0')
		}
//		rb = ringbuf.New(5)
		size := 5
		ringbuffer := make([]int8, size)
		head := 0
		maxProd := 0
		for b, err := br.ReadByte(); err == nil; b, err = br.ReadByte() {
			digit := toDigit(b)
			if digit < 0 {
				continue
			}
			ringbuffer[head] = digit
			head = (head + 1) % size
			curProd := 1
			for _, d := range ringbuffer {
				curProd *= int(d)
			}
			if curProd >= maxProd {
				maxProd = curProd
			}
		}
		return maxProd
	}},
	9: {1000, func(n int) int {
		// a^2 + b^2 = c^2
		// a + b + c = n
		// c = n - (a+b)
		// c^2 = n^2 + (a+b)^2 - 2*n*(a+b)
		// c^2 = n^2 + a^2 + b^2 + 2ab - 2n(a+b) = a^2 + b^2
		// (n^2)/2 + ab = n(a+b)
		// n/2 + ab/n = a+b
		// a*b is divisible by n
		if n%2 != 0 {
			fmt.Println("I'm not buying that! Sum of pythegoriean ints are even, fo shizzle!")
			return -1
		}
		for a := 1; a < n/2; a++ {
			// XXX this is not even right!
			for b := 1; b < a; b++ {
				// this is not really effective
				if a*b%n == 0 {
					if n/2+a*b/n == a+b {
						c := 1000 - a - b
						return a * b * c
					}
				}
			}
		}
		return -1
	}},
	10: {2000000, func(n int) int {
		// XXX copied from problem 7 but whatever
		primes := make([]int, 2)
		isPrime := func(k int) bool {
			to := int(math.Sqrt(float64(k)))
			for _, p := range primes {
				if p > to {
					return true
				}
				if k%p == 0 {
					return false
				}
			}
			return true
		}
		primes[0] = 2
		primes[1] = 3
		sum := 2
		to := int(math.Sqrt(float64(n)))
		for prevPrime := 3; prevPrime < n; {
			sum += prevPrime
			nextPrime := prevPrime + 2
			for ; !isPrime(nextPrime); nextPrime += 2 {
			}
			if nextPrime < to {
				// only need to store primes for isPrime
				primes = append(primes, nextPrime)
			}
			prevPrime = nextPrime
		}
		return sum
	}},
	11: {20, func(n int) int {
		br, closer := inputs.BufReader("inputs", 11, n)
		defer closer.Close()
		scanner := bufio.NewScanner(br)
		grid := make([][]uint8, n)
		for i:=0; scanner.Scan(); i++ {
			line := scanner.Text()
			grid[i] = make([]uint8, n)
			for j, valString := range strings.Split(line, " ") {
				val, _ := strconv.Atoi(valString)
				grid[i][j] = uint8(val)
			}
		}

		size := 4
		maxProd := 0
		for i := 0; i < n; i++ {
			ringbuffer := make([]uint8, size)
			head := 0
			for j := 0; j < n; j++ {
				ringbuffer[head] = grid[]
				head = (head + 1) % size
				curProd := 1
				for _, d := range ringbuffer {
					curProd *= int(d)
				}
				if curProd >= maxProd {
					maxProd = curProd
				}
			}
		} 

		return 0
	}},
}
