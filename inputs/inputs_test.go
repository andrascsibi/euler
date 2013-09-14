package inputs

import (
	"testing"
)

func TestToString(t *testing.T) {
	input, err := ToString(0, 0)
	if input != "hello\n" {
		t.Errorf("expected 0_0.txt to contain the string hello")
	}
	if err != nil {
		t.Errorf("expected err to be null")
	}
	input, err = ToString(0, 1)
	if err == nil {
		t.Errorf("expected some error as 0_1.txt does not exist")
	}
}
