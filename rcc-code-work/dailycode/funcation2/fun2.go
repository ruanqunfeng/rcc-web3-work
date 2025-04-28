package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

func reflectTest01(b interface{}) {
	// 通过反射获取到传入的变量的type， kind（类别），值
	// 1. 先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=", rTyp)

	// 2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	//n1 := 10
	//n2 := n1 + rVal.Interface().(int)
	// i := 10 + rVal.Interface().(int) + int(rVal.Int()) rVal.Int()返回的是int64类型
	i := 10 + int64(rVal.Interface().(int)) + rVal.Int()
	fmt.Println("rVal=", rVal)
	fmt.Println("i=", i)
}

type Student struct {
	Name string
	Age  int
}

func reflectTest02(b interface{}) {
	// 通过反射获取到传入的变量的type， kind（类别），值
	// 1. 先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=", rTyp)

	// 2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	kind1 := rTyp.Kind()
	kind2 := rVal.Kind()

	fmt.Printf("kind1=%v kind2=%v\n", kind1, kind2)

	iv := rVal.Interface()
	fmt.Printf("iv=%v iv=%T\n", iv, iv)
	// iv.Name 错误，因为反射是运行时识别到了iv是Student类型，但是编译阶段并不知道
	stu, ok := iv.(Student)
	if ok {
		fmt.Println(stu.Name, stu.Age)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "【监控1】")
	go watch(ctx, "【监控2】")
	go watch(ctx, "【监控3】")
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
