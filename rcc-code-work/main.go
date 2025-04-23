package main

import (
	"fmt"
)

type CustomError struct {
	code int
	msg  string
}

func (err *CustomError) Error() string {
	return fmt.Sprintf("%d - %s", err.code, err.msg)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &CustomError{arg, "can't work with it"}
		// return -1, CustomError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	_, e := f2(42)
	if ae, ok := e.(*CustomError); ok {
		fmt.Println(ae.code)
		fmt.Println(ae.msg)
	}
}
