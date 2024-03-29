package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.NewBuffer(nil)
	n, err := buf.Write([]byte("hello world"))
	if err != nil {
		fmt.Println("write error:", err)
	}
	fmt.Printf("write %d bytes\n", n) // 输出：write 11 bytes
	fmt.Println(buf.String())         // 输出：hello world
}
