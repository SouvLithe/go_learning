package main

import "fmt"

func main() {
	// 字面量赋值
	m := make(map[string]int, 10)
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m)
	// 字面量赋值
	m2 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(m2)

}
