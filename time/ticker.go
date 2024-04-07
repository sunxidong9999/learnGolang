package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个定时器，每隔 1 秒向通道发送一个时间事件
	ticker := time.NewTicker(1 * time.Second)

	// 启动一个 goroutine 来处理定时任务
	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				// 在这里可以执行需要周期执行的任务
			}
		}
	}()

	// 等待一段时间，然后停止定时器
	time.Sleep(5 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
