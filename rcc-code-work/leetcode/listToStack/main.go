package main

type MyStack struct {
	queue []int
	l     int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	if len(this.queue) > this.l {
		this.queue[this.l] = x
	} else {
		this.queue = append(this.queue, x)
	}
	this.l++
}

func (this *MyStack) Pop() int {
	x := this.queue[this.l-1]
	this.l--
	return x
}

func (this *MyStack) Top() int {
	return this.queue[this.l-1]
}

func (this *MyStack) Empty() bool {
	return this.l == 0
}
