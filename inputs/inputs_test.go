package inputs

import (
	"testing"
)

func TestString(t *testing.T) {
	input := String(".", 0, 0)
	if input != "hello\n" {
		t.Errorf("expected 0_0.txt to contain the string hello, but was: %v", input)
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("should have panicked")
		}
	}()
	input = String(".", 0, 1)
}

func TestBufReader(t *testing.T) {
	br, closer := BufReader(".", 0, 0)
	defer closer.Close()
	if in, _ := br.ReadString('\n'); in != "hello\n" {
		t.Errorf("expected 0_0.txt to contain the string hello, but was: %v", in)
	}
}
