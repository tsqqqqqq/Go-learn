package main

import (
	"fmt"
	"io/ioutil"
)

const (
	add = "+"
	sub = "-"
	mul = "*"
	div = "/"
)

func operation(operation string, a, b int) (res int) {
	switch operation {
	case add:
		res = a + b
	case sub:
		res = a - b
	case mul:
		res = a * b
	case div:
		res = a / b
	default:
		panic("不存在该操作符")
	}
	return
}

func main() {
	const filename = "abc.txt"
	contexts, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", contexts)

	res := operation(add, 3, 4)
	fmt.Println(res)
}
