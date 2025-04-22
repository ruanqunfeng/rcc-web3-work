package main

import (
	"fmt"
	"reflect"
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

func reflectTest03(b any) {
	v := reflect.ValueOf(b)
	k := v.Kind()

	switch k {
	case reflect.Int:
		fmt.Println("int类型")
	case reflect.Int8:
		fmt.Println("int8类型")
	case reflect.Chan:
		fmt.Println("chan类型")
	default:
		fmt.Println("其他类型")
	}
}

func main() {
	// var num int = 100
	// reflectTest01(num)

	// stu := Student{Name: "tome", Age: 18}
	// reflectTest02(stu)

	// reflectTest02(func(a int) func(int) int {
	// 	a++
	// 	return func(b int) int {
	// 		return a + b
	// 	}
	// })

	reflectTest03(100)
}
