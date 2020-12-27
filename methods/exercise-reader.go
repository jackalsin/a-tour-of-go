package main

import (
	"golang.org/x/tour/reader"
	errors "golang.org/x/xerrors"
)

type MyReader struct{}

func (MyReader) Read(b []byte) (int, error) {
	if b == nil {
		return 0, errors.New("buffer b cannot be nil")
	}
	l := len(b)
	for i := 0; i < l; i++ {
		b[i] = 'A'
	}
	return l, nil
}

func main() {
	reader.Validate(MyReader{})
}
