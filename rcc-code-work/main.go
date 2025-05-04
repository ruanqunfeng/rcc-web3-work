package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

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

var (
	myMap = make(map[int]uint, 200)
)

func test(n int, ch chan uint, exitChan chan bool) {
	sum := uint(1)
	for i := 1; i <= n; i++ {
		sum *= uint(i)
	}
	ch <- sum
	exitChan <- true
}

func main() {

	// c1 := make(chan uint, 49)
	// exitChan := make(chan bool, 1)

	// for i := 2; i <= 50; i++ {
	// 	go test(i, c1, exitChan)
	// }

	// for i := 2; i <= 50; i++ {
	// 	<-exitChan
	// }

	// close(c1)

	// for sum := range c1 {
	// 	fmt.Println(sum)
	// }

	m := make(map[any]any, 20)
	m[1] = 2

	arr := []int{1, 2, 3}
	m[arr] = 3

	fmt.Println(m)
}

func TestReflectFunc(t *testing.T) {
	call1 := func(v1, v2 int) {
		t.Log(v1, v2)
	}

	call2 := func(v1, v2 int, s string) {
		t.Log(v1, v2, s)
	}

	var (
		function reflect.Value
		inValue  []reflect.Value
		n        int
	)

	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)
		inValue = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			inValue[i] = reflect.ValueOf(args[i])
		}
		function = reflect.ValueOf(call)
		function.Call(inValue)
	}
	bridge(call1, 1, 2)
	bridge(call2, 1, 2, "test2")
}
