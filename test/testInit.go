package main

import (
	"bytes"
	"io"
)

const DEBUG = false

func main() {
	var buf *bytes.Buffer
	if DEBUG {
		buf = new(bytes.Buffer)
	}
}
func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("fdfdf"))
	}
}
