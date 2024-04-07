package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个每隔 1 秒发送一个时间事件的通道
	tick := time.Tick(1 * time.Second)

	// 从通道中接收时间事件，并输出
	for {
		select {
		case t := <-tick:
			fmt.Println("Tick at", t)
			// 在这里可以执行需要周期执行的任务
		}
	}
}
