package main

import (
	"fmt"
	"time"
)

func callBack() {
	fmt.Println("in AfterFunc() callback.")
}

func main() {
	fmt.Println("Start..")

	t := time.AfterFunc(1*time.Second, callBack)

	time.Sleep(3 * time.Second)

	t.Stop()

	fmt.Println("End")
}
