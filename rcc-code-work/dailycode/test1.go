package main

import (
	"fmt"
	"reflect"
	"time"
)

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{Name: "Easy", Age: 33}
	userType := reflect.TypeOf(user)
	fmt.Println(userType.Name()) // 输出结构体的名称
	fmt.Println(userType.Kind()) // 输出结构体的种类

	reflectValue := reflect.ValueOf(user)
	fmt.Println(reflectValue)

	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		value := reflectValue.Field(i).Interface()
		fmt.Printf("Field Name: %s, Field Type: %s, Field Value: %v\n", field.Name, field.Type, value)
	}

	u1 := reflect.New(userType).Elem() // 创建一个新的实例
	u1.FieldByName("Name").SetString("John")
	u1.FieldByName("Age").SetInt(25) // 设置字段值
	fmt.Println(u1.Interface())      // 输出: {John 25}

	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		// close(ch)
		fmt.Println("Channel closed in goroutine")
	}()

	time.Sleep(time.Second)

	for c := range ch {
		fmt.Println(c)
	}
}
