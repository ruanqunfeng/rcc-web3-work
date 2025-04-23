package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseBetween 翻转链表的指定区间
// 1 <= left <= right <= n
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{0, head}
	var lNode *ListNode = dummy
	var rNode *ListNode
	// 找到要翻转节点的前一个节点，可能是虚拟节点
	for i := 0; i < left-1; i++ {
		lNode = lNode.Next
	}

	// 找到要翻转的最后一个节点
	rNode = lNode
	for i := left; i <= right; i++ {
		rNode = rNode.Next
	}

	l1 := lNode.Next
	// l2是翻转链接表的下一个节点
	l2 := rNode.Next

	// 断开翻转的链表
	lNode.Next = nil
	rNode.Next = nil

	// 翻转链表
	// l1 是翻转的第一个节点，l2 是翻转的最后一个节点
	reverseList(l1)
	// 翻转之后rNode变成了第一个节点
	lNode.Next = rNode
	l1.Next = l2

	return dummy.Next

}

func reverseList(head *ListNode) {
	cur := head
	var r *ListNode
	for cur != nil {
		p := cur.Next
		cur.Next = r
		r = cur
		cur = p
	}
}

func main() {
	head := &ListNode{Val: 5}

	p := reverseBetween(head, 1, 1)

	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
