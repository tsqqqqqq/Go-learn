package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

type student struct {
	id    int
	names string
	sex   int
}

type class struct {
	id   int
	name string
	addr string
}

func (c *student) setId(n int) {
	c.id = n
}

/***
字符和字符串的处理
*/
func main() {
	str := "我喜欢你"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))
	for _, v := range []byte(str) {
		fmt.Printf("%X ", v)
	}
	fmt.Printf("%x", ^uintptr(0))

	stu := new(student)
	stu.names = "studentname"
	cla := new(class)
	cla.name = "classname"

	fmt.Println(reflect.TypeOf(stu)) //*main.student
	fmt.Println(reflect.TypeOf(cla)) //*main.class

	p := stu
	p = (*student)(unsafe.Pointer(cla))
	fmt.Println(reflect.TypeOf(p)) //*main.class
	fmt.Println(p.names)
	p.setId(10)
	fmt.Println(cla)
	fmt.Println(reflect.TypeOf(cla)) //*main.class
	fmt.Println(stu)
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(p)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[3:6:8]
	s2 := slice[3:6]
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
}
