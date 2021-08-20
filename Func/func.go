package main

import (
	"fmt"
	"reflect"
	"runtime"
)

const (
	add = "+"
	sub = "-"
	mul = "*"
	div = "/"
)

func operation(operation string, a, b int) (int, error) {
	res := 0
	var err error
	switch operation {
	case add:
		res, err = a+b, nil
	case sub:
		res, err = a-b, nil
	case mul:
		res, err = a*b, nil
	case div:
		res, err = a/b, nil
	default:
		fmt.Errorf("当前操作错误 导致错误的操作为 %s", operation)
	}
	return res, err
}

func call(operation func(op string, a, b int) (int, error), op string, a, b int) (int, error) {
	funcPtr := reflect.ValueOf(operation).Pointer()
	funcName := runtime.FuncForPC(funcPtr).Name()
	fmt.Printf("当前执行的函数名为 %s\n", funcName)
	return operation(op, a, b)
}

func main() {
	res, err := call(operation, add, 3, 4)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
