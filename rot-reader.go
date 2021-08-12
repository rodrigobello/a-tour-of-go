package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	if b >= 'a' && b <= 'm' || b >= 'A' && b <= 'M' {
		b += 13
	} else if b >= 'n' && b <= 'z' || b >= 'N' && b <= 'Z' {
		b -= 13
	}
	return b
}

func (rot13r *rot13Reader) Read(bytes []byte) (int, error) {
	_, err := rot13r.r.Read(bytes)
	if err != nil {
		return 0, err
	}
	for i := range bytes {
		bytes[i] = rot13(bytes[i])
	}
	return len(bytes), nil
}

func runRot13ReaderExercise() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
