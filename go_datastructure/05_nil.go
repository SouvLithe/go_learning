package main

import "fmt"

func main() {
	var a *int
	fmt.Println(a == nil)
	var b map[int]int
	fmt.Println(b == nil)

	// 无法比较类型不同的nil
	//fmt.Println(a == b)

	// nil永远不是一个结构体
	//var c struct{}
	//fmt.Println(c == nil)

	var d interface{}
	fmt.Println(d == nil)
	var c *int
	// a = d  // 这样赋值后d的eface字段的类型立马就有值了
	fmt.Println(c == nil)
	fmt.Println(a == nil)
}
