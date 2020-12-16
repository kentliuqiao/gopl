package e7_1

import (
	"fmt"
	"testing"
)

func TestWordCounter(t *testing.T) {
	pwc := new(WordCounter)

	n, err := fmt.Fprintf(pwc, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. ")

	if int(*pwc) != 19 || err != nil {
		t.Errorf("Failled to count: %d", n)
	}

	pwc.Write(nil)

	pwc.Write([]byte("Lorem ipsum"))
	if int(*pwc) != 21 {
		t.Errorf("Failled to accumulate")
	}
}

func TestLineCounter(t *testing.T) {
	plc := new(LineCounter)

	n, err := fmt.Fprintf(plc, "Lorem ipsum dolor sit amet\n consectetur adipiscing elit\n sed do eiusmod tempor incididunt ut labore et dolore magna aliqua\n")

	if int(*plc) != 3 || err != nil {
		t.Errorf("Failled to count: %d", n)
	}

	plc.Write(nil)

	plc.Write([]byte("Lorem\nipsum"))
	if int(*plc) != 5 {
		t.Errorf("Failled to accumulate")
	}
}
