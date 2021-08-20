package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
将整型转换为二进制字符串
*/
func convertToBin(n int) (res string) {
	//循环选择为 n/2
	for ; n > 0; n /= 2 {
		//获取每次循环值的低位
		lsb := n % 2
		//strconv.Itoa（n int） 将int型强制转换为string型 并且将字符依次向前加
		res = strconv.Itoa(lsb) + res
	}
	return
}

/**
循环读取文件内容
*/
func printFile(filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0325)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	//fmt.Println(convertToBin(0))
	printFile("./Branch/abc.txt")
}
