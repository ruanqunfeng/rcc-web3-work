package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	l := sync.Mutex{}
	l.Lock()
	defer l.Unlock()

	ws := sync.WaitGroup{}
	ws.Add(2)
	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go func(ctx context.Context) {
			i := 0
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("goroutine 2 exit")
					ws.Done()
					return
				default:
					fmt.Printf("goroutine 2: i = %d\n", i)
					time.Sleep(time.Second)
					i++
				}

			}
		}(ctx)

		for j := 0; j <= 3; j++ {
			fmt.Printf("goroutine 1: i = %d\n", j)
			time.Sleep(500 * time.Millisecond)
		}

		fmt.Println("goroutine 1 exit")
		ws.Done()
	}()

	ws.Wait()
	fmt.Println("Main exit")
}
