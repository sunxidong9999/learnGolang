package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.NewBufferString("hello world")
	data := make([]byte, 5)
	n, err := buf.Read(data)
	if err != nil {
		fmt.Println("read error:", err)
	}
	fmt.Printf("read %d bytes:\n", n) // 输出：read 5 bytes
	fmt.Println(string(data))         // 输出：hello
}
