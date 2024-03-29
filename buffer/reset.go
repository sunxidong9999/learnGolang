package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String()) // 输出：hello
	buf.Grow(60)
	fmt.Printf("len=%d, cap=%d\n", buf.Len(), buf.Cap()) // 输出：len=5, cap=64
	buf.Reset()
	fmt.Println(buf.String())                            // 输出：
	fmt.Printf("len=%d, cap=%d\n", buf.Len(), buf.Cap()) // 输出：len=5, cap=64
}
