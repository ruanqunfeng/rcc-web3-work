package main

import (
	"sync"
	"time"
)

// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。

func modifyValue(i *int) {
	*i += 10
}

func modifySlice(s *[]int) {
	for i := range *s {
		(*s)[i] *= 2
	}
}

// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。

var wg sync.WaitGroup

func printNum(b bool) {
	if b {
		for i := 1; i <= 10; i += 2 {
			println(i)
		}
	} else {
		for i := 2; i <= 10; i += 2 {
			println(i)
		}
	}
	wg.Done()
}

func printTask(task func()) {
	start := time.Now()
	task()
	elapsed := time.Since(start)
	println("Task executed in:", elapsed)
	wg.Done()
}

func main() {
	var num int = 5
	modifyValue(&num)
	println(num) // 输出：15

	slice := []int{1, 2, 3}
	modifySlice(&slice)
	for _, v := range slice {
		println(v) // 输出：2 4 6
	}

	wg.Add(2)
	go printNum(true)
	go printNum(false)

	wg.Add(1)
	go printTask(func() {
		time.Sleep(2 * time.Second)
		println("Task 1 completed")
	})

	wg.Wait()
}
