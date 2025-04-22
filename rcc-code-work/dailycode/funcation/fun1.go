package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	initCh := make(chan int, 1000)
	permCh := make(chan int, 1000)
	exitCh := make(chan bool, 1)

	//wg.Add(1)
	go func() {
		for i := range 120000 {
			initCh <- i
		}
		close(initCh)
	}()

	for range 16 {
		go func(initCh, permCh chan int, exitchan chan bool) {
			for i := range initCh {
				flag := true
				// 计算素数
				for j := 2; j < i; j++ {
					if i%j == 0 {
						flag = false
						break
					}
				}

				if flag {
					permCh <- i
				}

			}

			exitchan <- true
		}(initCh, permCh, exitCh)
	}

	wg.Add(1)
	go func(chan int) {
		for i := range permCh {
			i++
		}
		wg.Done()
	}(permCh)

	wg.Add(1)
	go func(chan bool) {
		for range 16 {
			<-exitCh
		}
		close(permCh)
		wg.Done()
	}(exitCh)

	wg.Wait()
	fmt.Println("main out")

}
