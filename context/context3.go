package main

import (
	"context"
	"log"
	"os"
	"time"
)

// 通过context追踪日志

func main() {
	// 创建一个 context 并设置日志记录
	ctx := context.WithValue(context.Background(), "logger", log.New(os.Stdout, "", 0))

	// 在另一个 goroutine 中执行任务
	go func(ctx context.Context) {
		// 获取日志记录器
		logger := ctx.Value("logger").(*log.Logger)

		// 记录日志
		logger.Println("执行任务")
	}(ctx)

	// 等待另一个 goroutine 结束
	time.Sleep(3 * time.Second)
}
