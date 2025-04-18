package main

import (
	"fmt"
	gobasework "rcc-code-work/go-base-work"
)

func main() {
	nums := [][]int{{1, 4}, {2, 3}, {8, 10}, {15, 18}}
	num := gobasework.Merge(nums)
	fmt.Println(num)
}
