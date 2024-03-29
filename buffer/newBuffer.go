package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 方法1
	buf1 := bytes.NewBufferString("hello world")
	fmt.Println(buf1.String()) // 输出：hello world

	// 方法2
	var buf bytes.Buffer
	buf.WriteString("hello")
	buf.WriteString(" ")
	buf.WriteString("world")
	fmt.Println(buf.String()) // 输出：hello world
}
