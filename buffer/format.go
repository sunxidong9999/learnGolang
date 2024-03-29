package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	//buf := bytes.NewBufferString("hello")
	fmt.Printf("len = %d, cap = %d\n", buf.Len(), buf.Cap())

	//	buf.Grow(60)
	//	fmt.Printf("len=%d, cap=%d\n", buf.Len(), buf.Cap()) // 输出：len=5, cap=16

	for i := 0; i < 10; i++ {
		fmt.Fprintf(&buf, "%d", i)
	}
	fmt.Println(buf.String())
	fmt.Printf("len = %d, cap = %d\n", buf.Len(), buf.Cap())
}
