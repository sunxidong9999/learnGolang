package main

import (
	"context"
	"fmt"
	"time"
)

// 通过context在协程间传递数据

func main() {
	// 创建一个 context
	ctx := context.Background()

	// 将数据存储到 context 中
	ctx = context.WithValue(ctx, "key", "value")

	// 在另一个 goroutine 中获取数据
	go func(ctx context.Context) {
		// 从 context 中获取数据
		value := ctx.Value("key")
		fmt.Println("value of key:", value) // 输出: value
		ctx.Done()
	}(ctx)

	// 等待另一个 goroutine 结束
	time.Sleep(time.Second)
}
