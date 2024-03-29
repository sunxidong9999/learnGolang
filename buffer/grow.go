package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.NewBufferString("hello")
	//buf := bytes.NewBufferString("")
	fmt.Printf("len=%d, cap=%d\n", buf.Len(), buf.Cap()) // 输出：len=5, cap=32
	buf.Grow(60)
	fmt.Printf("len=%d, cap=%d\n", buf.Len(), buf.Cap()) // 输出：len=5, cap=16
}
