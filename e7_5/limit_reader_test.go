package e7_5

import (
	"io"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	r := LimitReader(strings.NewReader("ABCDE"), 3)
	p := make([]byte, 2)

	n, err := r.Read(p)
	if n != 2 || err != nil {
		t.Errorf("Unexpected read size")
	}

	total := n
	n, err = r.Read(p)
	total += n
	if total != 3 || err != nil {
		t.Errorf("Off limits read. total: %d, err: %v", total, err)
	}

	n, err = r.Read(p)
	if n != 0 || err != io.EOF {
		t.Errorf("Off limits read")
	}
}
