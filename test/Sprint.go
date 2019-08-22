package main

import (
	"io"
	"os"
	"strconv"
)

func Sprint(x interface{}) string {
	type stirnger interface {
		String() string
	}
	switch x := x.(type) {
	case stirnger:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		} else {
			return "false"
		}
	default:
		return "????"

	}
}

func main()  {
	var w io.Writer = os.Stdout
}
