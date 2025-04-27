package main

import (
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/rpc"
)

type A struct {
	i int
}

func (a *A) add(v int) int {
	a.i += v
	return a.i
}

// 声明函数变量
var funcation1 func(int) int

// 声明闭包
var squart2 func(int) int = func(p int) int {
	p *= p
	return p
}

func ParseBLockNumber(s string) (int64, error) {
	var err error
	var number int64
	switch s {
	case "head":
		number = int64(rpc.LatestBlockNumber)
	case "finalized":
		number = int64(rpc.FinalizedBlockNumber)
	case "safe":
		number = int64(rpc.SafeBlockNumber)
	default:
		number, err = strconv.ParseInt(s, 10, 64)
	}

	return number, err
}

func main() {

	s := "0xde2395eddf141f9592c5efff421bf5628b3246ab3680d1d74eaccae6d226588e"
	fmt.Println(len(s))

}
