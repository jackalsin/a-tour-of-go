package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(b []byte) (n int, err error) {
	n, err = reader.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

func rot13(b byte) byte {
	if 'a' <= b && b <= 'z' {
		return rot13Helper('a', 'z', b)
	} else if 'A' <= b && b <= 'Z' {
		return rot13Helper('A', 'Z', b)
	} else {
		return b
	}
}

func rot13Helper(l byte, h byte, b byte) byte {
	bias := b - l
	if bias < 13 {
		return b + 13
	}
	return b - 13
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
