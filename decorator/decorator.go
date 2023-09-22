package decorator

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func Decorate(r io.Reader, border string) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)

	fmt.Println(strings.Repeat(border, 20))
	fmt.Println(buf.String())
	fmt.Println(strings.Repeat(border, 20))
}
