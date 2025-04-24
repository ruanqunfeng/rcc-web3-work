package main

import "fmt"

// 分隔链表

// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

// 你应当 保留 两个分区中每个节点的初始相对位置。

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}

	tail1 := tail

	dummy := &ListNode{Val: 0, Next: head}
	dummy1 := dummy
	for dummy1 != nil && dummy1.Next != nil {
		// 防止无限循环, 因为dummy1没有每次都移动
		flag := false
		if dummy1.Next == tail {
			flag = true
		}
		// dummy1.Next.Next != nil 规避[1,2] 2这种情况,因为2已经是最后了,没必要再往后移动
		if dummy1.Next.Val >= x && dummy1.Next.Next != nil {
			next := dummy1.Next
			dummy1.Next = next.Next
			next.Next = nil
			tail1.Next = next
			tail1 = tail1.Next
		} else {
			dummy1 = dummy1.Next
		}
		if flag {
			break
		}
	}

	return dummy.Next

}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	// head.Next.Next = &ListNode{Val: 3}
	// head.Next.Next.Next = &ListNode{Val: 2}
	// head.Next.Next.Next.Next = &ListNode{Val: 5}
	// head.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	p := partition(head, 2)
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
