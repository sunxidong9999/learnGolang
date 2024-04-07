package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		ch <- "hello from channel"
	}()

	msg := <-ch

	fmt.Println(msg)

	// 缓冲通道(Buffered Channels)
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	ch2 <- 3
	ch2 <- 5
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	ch3 := make(chan string)
	ch4 := make(chan string)
	go func() {
		time.Sleep(time.Second)
		ch3 <- "hello from channel3"
	}()
	go func() {
		time.Sleep(time.Second)
		ch4 <- "hello from channel4"
	}()

	// select语句会阻塞直到其中一个case可以继续进行，此时就执行该case。
	// 也就是只会执行其中的一个case。
	select {
	case msg3 := <-ch3:
		fmt.Println(msg3)
	case msg4 := <-ch4:
		fmt.Println(msg4)
	// 设置等待超时时间
	case <-time.After(2 * time.Second): // 设置超时时间为 1 秒
		fmt.Println("timeout")

	}

	close(ch)
	close(ch2)
	close(ch3)
	close(ch4)
}
