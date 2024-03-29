package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// 读取字节数组
	data := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f}
	var buf bytes.Buffer
	buf.Write(data)

	// 转换大小端序
	var num uint16
	binary.Read(&buf, binary.BigEndian, &num)
	fmt.Println(num) // 输出：0x4865

	// 写入字节数组
	data2 := []byte{0x57, 0x6f, 0x72, 0x6c, 0x64, 0x21}
	buf.Write(data2)
	fmt.Println(buf.Bytes()) // 输出：[72 101 108 108 111 87 111 114 108 100 33]
}
