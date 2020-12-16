package e7_2

import (
	"io"
)

type countWriter struct {
	w io.Writer
	c int64
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &countWriter{w: w, c: 0}

	return cw, &(cw.c)
}

func (cw *countWriter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)

	cw.c += int64(n)
	return n, err
}
