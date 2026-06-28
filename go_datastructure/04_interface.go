package main

import "fmt"

type Car interface {
	Drive()
}

type Truck struct {
	Model string
}

func (t Truck) Drive() {
	fmt.Println("t.Model = ", t.Model)
}

type empty interface {
}

func K(interface{}) {

}

func main() {
	var c Car = Truck{}
	fmt.Println(c)
	t := c.(Truck)
	fmt.Println(t)

	switch c.(type) {
	case Truck:
		fmt.Println("Truck")
	case Car:
		fmt.Println("Car")
	}

	// 可以实现类似泛型的效果
	K(123)
	K("nihao")
}
