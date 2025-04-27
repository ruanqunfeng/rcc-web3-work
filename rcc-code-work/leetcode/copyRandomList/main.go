package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 随机链表的复制
var copyNode map[*Node]*Node

func copyRandomList(head *Node) *Node {
	copyNode = make(map[*Node]*Node)
	return deepCopy(head)
}

func deepCopy(n *Node) *Node {
	if n == nil {
		return nil
	}

	if v, ok := copyNode[n]; ok {
		return v
	}

	newNode := &Node{Val: n.Val}
	copyNode[n] = newNode
	newNode.Next = deepCopy(n.Next)
	newNode.Random = deepCopy(n.Random)
	return newNode
}

func main() {

}
