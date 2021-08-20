package main

import (
	"fmt"
	"math"
)

/**
go语言的常量关键字为consts 使用方法与变量var一致
特殊的是，常量在声明的时候不定义具体类型，该常量则为一个文本，即可以是int 也可以是float
*/
func consts() {
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

/**
自定义枚举类型
*/
const (
	cpp    = "cpp"
	java   = "java"
	python = "python"
)

const (
	b = 1 << (10 * iota)
	kb
	mb
	gb
	tb
	pb
)

func main() {
	consts()

	fmt.Println(cpp, java, python)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
