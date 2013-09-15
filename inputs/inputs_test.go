package inputs

import (
	"testing"
)

func TestToString(t *testing.T) {
	input := ToString(".", 0, 0)
	if input != "hello\n" {
		t.Errorf("expected 0_0.txt to contain the string hello")
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("should have panicked")
		}
	}()
	input = ToString(".", 0, 1)
}
