package fabnocci

import (
	"fmt"
	"io"
	"strings"
)

// 一般实现
func Fabnocci() Fab {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type Fab func() int

func (f *Fab) Next() int {
	return (*f)()
}

func (fa Fab) Read(p []byte) (n int, err error) {
	next := fa()
	if next > 10000000000000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}
