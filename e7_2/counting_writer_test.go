package e7_2_test

import (
	"bytes"
	"testing"

	"gopl/e7_2"
)

func TestCountingWriter(t *testing.T) {
	buf := new(bytes.Buffer)
	w, nptr := e7_2.CountingWriter(buf)

	s := "1234567890"
	w.Write([]byte(s))
	if *nptr != 10 {
		t.Errorf("Number of bytes counted should be 10: %d", *nptr)
	}
	if buf.String() != "1234567890" {
		t.Errorf("Buffer shoud be equals to '%s'", s)
	}

	w.Write([]byte(""))
	if *nptr != 10 {
		t.Errorf("Number of bytes counted should be 10: %d", *nptr)
	}

	w.Write([]byte("123"))
	if *nptr != 13 {
		t.Errorf("Number of bytes counted should be 13: %d", *nptr)
	}

	w.Write([]byte(nil))
	if *nptr != 13 {
		t.Errorf("Number of bytes counted should be 13: %d", *nptr)
	}

}
