package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个 context 并设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	// 在另一个 goroutine 中执行任务
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				// 超时或取消
				fmt.Println("超时或取消")
				return
			default:
				// 执行任务
				fmt.Println("执行任务")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	// 等待 2 秒
	time.Sleep(2 * time.Second)

	// 取消 goroutine
	cancel()

	// 等待另一个 goroutine 结束
	// 等待 2 秒
	time.Sleep(2 * time.Second)
}
