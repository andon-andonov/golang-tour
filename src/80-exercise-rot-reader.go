package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, e := r.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, e
}

func rot13(b byte) byte {
	if b >= 65 && b <= 90 {
		return ((b - 65 + 13) % 26) + 65
	}
	if b >= 97 && b <= 122 {
		return ((b - 97 + 13) % 26) + 97
	}
	return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
