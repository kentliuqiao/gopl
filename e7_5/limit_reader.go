package e7_5

import (
	"io"
)

type limitReader struct {
	reader io.Reader
	limit  int64
	pos    int64
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{reader: r, limit: n}
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	var readLen int64

	if lr.pos == lr.limit && len(p) > 0 {
		return 0, io.EOF
	}

	if lr.limit < lr.pos+(int64(len(p))) {
		readLen = lr.limit - lr.pos
	} else {
		readLen = int64(len(p))
	}

	n, err = lr.reader.Read(p[:readLen])
	lr.pos += int64(n)

	return
}
