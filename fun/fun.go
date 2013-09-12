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
