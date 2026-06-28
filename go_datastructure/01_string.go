package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Sizeof("nihao woshishei"))
	s := "hello world"
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(sh.Len)
	for i := 0; i < len(s); i++ {
		// 可以发现这种方式访问的是底层的字节
		fmt.Printf("%x ", s[i])
	}

	// 使用range进行遍历的时候， 会被解码成rune类型的字符
	//解码算法位于 runtime/utf8.go文件中
	for _, char := range s {
		// 这种访问的才是数组的元素，第一个是角标，第二个才是元素
		fmt.Printf("\n%c ", char)
	}
}
