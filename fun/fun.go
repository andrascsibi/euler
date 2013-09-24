package fun

import (
	"fmt"
)

// Alright, this is not functional at all,
// I just couldn't figure out where to put it
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

func All(list []int, predicate func(int) bool) bool {
	for _, item := range list {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func Seq(from, to int) []int {
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

func LazyAll(list chan int, predicate func(int) bool) bool {
	for item := range list {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func LazySeq(from, to int) chan int {
	c := make(chan int)
	go func() {
		for i := from; i < to; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

type Reducer func(cur, acc int) int

func Reduce(slice []int, reducer Reducer, init int) int {
	acc := init
	for _, elem := range slice {
		acc = reducer(elem, acc)
	}
	return acc
}

func Sum(a, b int) int {
	return a + b
}

func Prod(a, b int) int {
	return a * b
}

type Mapper func(n int) int

func Memoize(memo []int, mapper Mapper) Mapper {
	return func(n int) int {
		ans := 0
		if n < len(memo) {
			ans = memo[n]
			if ans != 0 {
				fmt.Println(ans)
				return ans
			}
		}
		ans = mapper(n)
		if n < len(memo) {
			memo[n] = ans
		}
		return ans
	}
}
