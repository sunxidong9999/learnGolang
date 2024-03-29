package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.NewBufferString("hello world")
	buf.Truncate(8)
	fmt.Println(buf.String()) // 输出：hello
}
