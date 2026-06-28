package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	fmt.Println(s)
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := arr[1:4]
	fmt.Printf("len=%d, cap=%d\n", len(slice), cap(slice))
	ints := append(s, 4, 5, 6)
	fmt.Printf("len=%d, cap=%d\n", len(ints), cap(ints))
}
