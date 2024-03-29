package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 从文件中读取数据
	file, err := os.Open("encoder.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	if err != nil {
		log.Fatal(err)
	}
	// 将数据输出到终端
	fmt.Println(buf.String())

	// 将数据写入文件
	out, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, &buf)
	if err != nil {
		log.Fatal(err)
	}
}
