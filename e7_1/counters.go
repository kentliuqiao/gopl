package e7_1

import (
	"bufio"
	"bytes"
	"io"
)

type WordCounter int

type LineCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	n, err := splitCounter(bytes.NewBuffer(p), bufio.ScanWords)
	if err != nil {
		return 0, err
	}
	*w += WordCounter(n)
	return n, nil
}

func (w *LineCounter) Write(p []byte) (int, error) {
	n, err := splitCounter(bytes.NewBuffer(p), bufio.ScanLines)
	if err != nil {
		return 0, err
	}
	*w += LineCounter(n)
	return n, nil
}

func splitCounter(r io.Reader, split bufio.SplitFunc) (int, error) {
	count := 0
	scanner := bufio.NewScanner(r)
	scanner.Split(split)
	for scanner.Scan() {
		count++
	}

	return count, scanner.Err()
}
