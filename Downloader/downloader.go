package main

import (
	"fmt"
	"learngo/infar"
)

func main() {
	rf := infar.RetrieveFactory{}
	fmt.Println(rf.GetRetrieve(infar.TESTRETRIEVE).Get("http://www.tsqqqqqq.cn"))

	fmt.Println(rf.GetTypeAssertionRetrieve(infar.TestRetrieve{"www.baidu.com"}).Get("www.baidu.com"))

	//r :=*(*infar.Retrieve)(unsafe.Pointer(&infar.TestRetrieve{"www.baidu.com"}))
	//r.Get("string")
}
