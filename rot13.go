package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(sb []byte) (int, error) {
	s, err := r.r.Read(sb)
	for i := 0; i < s; i++ {
		if sb[i] > 64 && sb[i] < 91 {
			sb[i] = ((sb[i] + 13 - 65) % 26) + 65
		}
		if sb[i] > 96 && sb[i] < 123 {
			sb[i] = ((sb[i] + 13 - 97) % 26) + 97
		}

	}
	fmt.Println(string(sb))
	return s, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
