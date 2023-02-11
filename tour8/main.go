package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)
	if err != nil {
		return n, err
	}
	for i, v := range b {
		switch {
		case 'A' <= v && v < 'N', 'a' <= v && v < 'n':
			b[i] += 13
		case 'N' <= v && v <= 'Z', 'n' <= v && v <= 'z':
			b[i] -= 13
		}
	}
	return n, nil
}

// Exercise: rot13Reader
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
