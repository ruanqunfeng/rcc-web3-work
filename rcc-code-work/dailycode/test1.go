package main

import (
	"fmt"
	"sync"
	"time"
	_ "time"
)

var w sync.WaitGroup

func test1() {
	for i := range 10 {
		fmt.Println("test1", i)
	}
	w.Done()
}

func test2() {
	for i := range 10 {
		fmt.Println("test2", i)
	}
	w.Done()
}

func main() {
	w.Add(1)
	go test1()

	w.Add(1)
	go test2()

	w.Wait()

	fmt.Println("main out")

	ch := make(chan int)
	go func() {
		for {
			time.Sleep(1 * time.Second) // 未阻塞，周期性唤醒
		}
	}()
	<-ch // 主 goroutine 阻塞
}
