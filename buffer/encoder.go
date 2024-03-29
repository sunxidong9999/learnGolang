package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var buf bytes.Buffer

	// 将结构体编码为 JSON
	p := Person{"Alice", 25}
	enc := json.NewEncoder(&buf)
	enc.Encode(p)
	fmt.Println(buf.String()) // 输出：{"Name":"Alice","Age":25}

	// 从 JSON 解码为结构体
	var p2 Person
	dec := json.NewDecoder(&buf)
	dec.Decode(&p2)
	fmt.Printf("Name: %s, Age: %d\n", p2.Name, p2.Age) // 输出：Name: Alice, Age: 25
}
