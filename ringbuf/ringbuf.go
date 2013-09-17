package ringbuf

import (
	"github.com/andrascsibi/euler/fun"
)

type RingBuffer struct {
	buffer []int
	head   int
	size   int
}

func New(size int) RingBuffer {
	return RingBuffer{make([]int, size), 0, size}
}

func (rb *RingBuffer) Add(elem int) {
	rb.buffer[rb.head] = elem
	rb.head = (rb.head + 1) % rb.size
}

func (rb *RingBuffer) Reduce(reducer fun.Reducer, init int) int {
	return fun.Reduce(rb.buffer, reducer, init)
}
