package main

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2

	}
}

func main() {
	time1 := time.NewTicker(time.Millisecond * 500)
	go func() {
		// 打点器和定时器的机制有点相似：一个通道用来发送数据。这里我们在这个通道上使用内置的 range 来迭代值每隔500ms 发送一次的值。
		for t := range time1.C {
			fmt.Println("Tick at", t)
			time.Sleep(time.Second * 2)
		}
	}()

	for {
	}
}
