package main

import (
	"fmt"
)

/**
map 的key使用hash表 ，
slice map function 不可以作为key
Struct类型不包含上述字段也可以作为key
*/
func main() {
	//create map
	m1 := map[string]string{"name": "tsq", "age": "18"}

	m2 := make(map[string]string) //empty 开辟了内存空间

	var m3 map[string]string //nil 为nil 可以做运算

	fmt.Println(m1, m2, m3)

	//loop 无序
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//需要手动的先把key放到slice里排序，根据排序规则，
	//再遍历slice中的key，获取map

	//get Value
	name := m1["name"]
	fmt.Println(name)

	//check key
	sex, ok := m1["sex"]
	if ok {
		fmt.Println(sex)
	} else {
		fmt.Println("不存在当前key")
	}

	//delete values
	delete(m1, "sex")
	findStr("abcabcbb")
}

func findStr(str string) {
	maps := make(map[rune]int)
	start := 0
	max := 0
	for index, value := range str {
		if ch, ok := maps[value]; ok && ch >= start {
			start = ch + 1
		}
		if index-start+1 > max {
			max = index - start + 1
		}
		maps[value] = index
	}
	fmt.Println(max)
	fmt.Println(maps)
	strs := str[start:max]
	fmt.Println(strs)
}
