package main

import "fmt"

/**
这里使用的是值传递，即将参数的值拷贝一份到该函数中，并不会改变原函数的值
*/
func f1(a int) {
	a++
}

/**
这里传递的是变量a的指针，它与man函数的变量a同时指向同一个内存空间，即会改变原来的值
*/
func f2(a *int) {
	*a++
}

func main() {
	var a = 3
	f1(a)
	fmt.Println(a)
	f2(&a)
	fmt.Println(a)
}
