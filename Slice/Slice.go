package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[2:6]
	slice := make([]int, 3, 10)
	slice = append(slice, slice...)
	fmt.Println("update s1:", s1)
	s2 := s1[3:5]
	fmt.Println("after s1:", s1)
	fmt.Println(s2)
	s3 := append(s2, 10)
	fmt.Printf("%v %p\n", s3, &s3)
	s4 := append(s3, 11)
	fmt.Printf("%v %p\n", s4, &s4)
	s5 := append(s4, 12)
	fmt.Printf("%v %p\n", s5, &s5)
	fmt.Println(s2)
	fmt.Println(arr)

}
