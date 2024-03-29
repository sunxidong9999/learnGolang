package main

import (
	"bytes"
	"fmt"
)

func concatStrings(strs ...string) string {
	var buf bytes.Buffer
	for _, s := range strs {
		buf.WriteString(s)
	}
	return buf.String()
}

func main() {
	s1 := "hello"
	s2 := "world"
	s3 := "!"
	s := concatStrings(s1, " ", s2, s3)
	fmt.Println(s) // 输出：hello world!
}
