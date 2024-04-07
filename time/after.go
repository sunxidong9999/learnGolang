package main

import (
	"fmt"
	"time"
)

// code from https://mp.weixin.qq.com/s/itHsebcV7a2G5Wy87bW63A
func main() {
	tick := time.Tick(1e8)
	boom := time.After(7e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println(" .")
			time.Sleep(5e7)
		}
	}
}
