package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// 作用：1.随机sleep一会；2.如果入参ch不为空，会把sleep的时间给到ch
func sleepRandom(fromFunction string, ch chan int) {
	defer func() {
		fmt.Println(fromFunction, "sleepRandom complete")
	}()
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	randomNumber := r.Intn(100)
	sleeptime := randomNumber + 100
	fmt.Println(fromFunction, "Starting sleep for", sleeptime, "ms")
	time.Sleep(time.Duration(sleeptime) * time.Millisecond)
	fmt.Println(fromFunction, "Waking up, slept for ", sleeptime, "ms")
	if ch != nil {
		ch <- sleeptime
	}
}
func sleepRandomContext(ctx context.Context, ch chan bool) {
	defer func() {
		fmt.Println("sleepRandomContext complete")
		// 通过channel，通知上游执行完毕
		ch <- true
	}()
	sleeptimeChan := make(chan int)
	// 开启新的协程G2，让该协程执行逻辑，执行完毕后，通过sleeptimeChan通知执行完毕
	go sleepRandom("sleepRandomContext", sleeptimeChan)
	select {
	case <-ctx.Done():
		// 场景1：main()调用cancelFunction()
		// 场景2：doWorkContext()调用cancelFunction()
		// 场景3：doWorkContext()自动超时
		fmt.Println("sleepRandomContext: Time to return")
	case sleeptime := <-sleeptimeChan:
		// 当新的协程G2执行完毕，调用ch<-sleeptime时
		fmt.Println("Slept for ", sleeptime, "ms")
	}
}
func doWorkContext(ctx context.Context) {
	// 生成新的ctx，超时时间为150ms
	ctxWithTimeout, cancelFunction := context.WithTimeout(ctx, time.Duration(150)*time.Millisecond)
	defer func() {
		fmt.Println("doWorkContext complete")
		// 下游所有的ctx都会关闭
		cancelFunction()
	}()
	ch := make(chan bool)
	// 启动新的协程G1
	go sleepRandomContext(ctxWithTimeout, ch)
	select {
	case <-ctx.Done():
		// 当main退出，调用main的cancelFunction()时
		fmt.Println("doWorkContext: Time to return")
	case <-ch:
		// 当新的协程G1退出，执行ch<-true时
		fmt.Println("sleepRandomContext returned")
	}
}

func main() {
	ctx := context.Background()
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)
	defer func() {
		fmt.Println("Main Defer: canceling context")
		// 下游所有的ctx都会关闭
		cancelFunction()
	}()
	go func() {
		// main函数sleep一会
		sleepRandom("Main", nil)
		// 下游所有的ctx都会关闭
		cancelFunction()
		fmt.Println("Main Sleep complete. canceling context")
	}()
	doWorkContext(ctxWithCancel)
}
