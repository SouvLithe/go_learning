package main

import (
	"fmt"
	"unsafe"
)

type Empty struct{}
type Empty2 struct{}

func main() {
	a := Empty{}
	b := Empty2{}
	fmt.Println(unsafe.Sizeof(int(0)))
	fmt.Println(unsafe.Sizeof(int32(0)))
	fmt.Println(unsafe.Sizeof(int64(0)))
	fmt.Println(unsafe.Sizeof(uint(0)))
	fmt.Println(unsafe.Sizeof(a))
	fmt.Printf("%p =? %p", &a, &b)

	m := map[string]struct{}{} // hashset
	m["a"] = struct{}{}
	c := make(chan struct{})
	fmt.Printf("\n%p", c)
}
