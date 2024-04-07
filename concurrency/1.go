package main

import (
	"fmt"
	"time"
)

func sum(idx int, a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}

	fmt.Println(idx, time.Now())
	c <- total // 把结果发送到channel c
	fmt.Println(idx, time.Now())
}

func test1() {

	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	// 用协程计算前一半的值
	go sum(1, a[:len(a)/2], c)

	// 用协程计算后一半的值
	go sum(2, a[len(a)/2:], c)

	x, y := <-c, <-c // receive from c

	fmt.Println(1, x, y, x+y)

	close(c)
}

func test2() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	// 用协程计算后一半的值
	go sum(3, a[len(a)/2:], c)

	// 用协程计算前一半的值
	go sum(4, a[:len(a)/2], c)

	x := <-c
	y := <-c

	fmt.Println(2, x, y, x+y)

	close(c)
}

func main() {
	go test1()
	go test2()

	// 缓冲通道(Buffered Channels)
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2

	go func() {
		ch2 <- 3
	}()
	go func() {
		ch2 <- 4
	}()
	go func() {
		ch2 <- 5
	}()
	go func() {
		ch2 <- 6
	}()
	go func() {
		ch2 <- 7
	}()

	go func() {
		time.Sleep(time.Second)
		close(ch2)
	}()

	for {
		val, ok := <-ch2
		if !ok {
			// 通道已关闭并且为空，退出循环
			fmt.Println("bye")
			break
		}
		fmt.Println("Received:", val)
	}
}
