package e7_4

import (
	"io"
)

type stringReader struct {
	pos int
	str string
}

func NewReader(str string) io.Reader {
	return &stringReader{str: str, pos: 0}
}

func (sr *stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, sr.str[sr.pos:])

	sr.pos += n

	if n == 0 && sr.pos == len(sr.str) {
		err = io.EOF
	}
	return
}
