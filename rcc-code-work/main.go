package main

import (
	"cmp"
	"context"
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"testing"
	"time"

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

	// ch := make(chan struct{})

	// go func() {
	// 	close(ch)
	// }()

	// time.Sleep(time.Second)
	// <-ch
	// <-ch
	// <-ch
	// fmt.Println("main done")

	// s := []int{1, 2, 3}
	// for _, v := range s {
	// 	s = append(s, v)
	// 	fmt.Printf("len(s)=%v\n", len(s))
	// }
	// m := make(map[int]any)
	// m[1] = 1
	// m[2] = 2
	// m[3] = 3
	// keys := SortMapByKey(m)
	// for _, k := range keys {
	// 	fmt.Println(k)
	// }

	// fmt.Println(runtime.NumCPU())

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Work done")
	case <-ctx.Done():
		fmt.Println("Timeout:", ctx.Err()) // 输出: context deadline exceeded
	}

	ctx = context.WithValue(context.Background(), "user", "Alice")
	if user, ok := ctx.Value("user").(string); ok {
		fmt.Println("User:", user) // 输出: User: Alice
	}
}

// T 必须是可排序的类型（如 int、string 等）。
func SortMapByKey[T cmp.Ordered](mp map[T]any) []T {
	arr := make([]T, 0, len(mp))
	for k := range mp {
		arr = append(arr, k)
	}
	slices.Sort(arr)
	return arr
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
