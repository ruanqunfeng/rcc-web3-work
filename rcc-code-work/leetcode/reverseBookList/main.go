package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBookList(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	var res []int = make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

func main() {
	
}
