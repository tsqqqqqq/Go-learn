package main

import (
	"fmt"
)

/**
go 看似为全局变量，实际并非全局变量，其作用域仅在同一个包环境下。
*/
var aa = 12

/**
可以使用var()对变量集中定义
*/
var (
	ss = "abc"
	bb = true
)

/**
变量值初始化
语法：关键字 变量名 变量类型
例如：var name string
*/
func variableZeroValue() {
	// 分别创建变量age，name 变量类型分别为 int 、string型
	var age int
	var name string
	fmt.Println(age, name)
}

/**
变量赋初值
*/
func variableInitiaValue() {
	// go语言支持同时声明多个变量赋初值
	var age, sex = 18, 1
	var name = "tsqqqqqq"
	fmt.Println(age, sex, name)
}

/**
go 语言支持变量推断，即不声明变量类型，但是对变量进行初始化就会自动推断变量类型
*/
func variableTypeDeduction() {
	var a, b = 12, "abc"
	fmt.Println(a, b)
}

/**
go语言支持简易声明变量  使用语法糖 :=对变量进行初始化
但该语法糖只可以在函数内使用，无法在创建包范围内的变量时使用
*/
func variableShorter() {
	a, b := 12, true
	fmt.Println(a, b)
}

func main() {
	variableZeroValue()
	variableInitiaValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, ss)

}
