package gobasework

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

// Flatten 递归的方式将多级双向链表展开成单级双向链表
// 复杂度分析：时间复杂度O(n)，空间复杂度O(1)
// 其中n为链表的节点数

func Flatten(root *Node) *Node {
	dfs(root)
	return root
}

// last 记录当前链表的最后一个节点
func dfs(root *Node) (last *Node) {
	cur := root
	for cur != nil {
		next := cur.Next
		child := cur.Child
		if child != nil {
			childLast := dfs(child)

			// 连接当前节点和子链表的头节点
			cur.Next = cur.Child
			cur.Child.Prev = cur

			// 连接子链表的尾节点和当前节点的下一个节点
			if next != nil {
				childLast.Next = next
				next.Prev = childLast
			}

			cur.Child = nil
			last = childLast
		} else {
			last = cur
		}

		cur = next
	}

	return
}
