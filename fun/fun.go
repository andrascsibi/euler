package fun

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
