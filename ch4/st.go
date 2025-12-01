package main

import "fmt"

// Human 结构体
type Human struct {
	Age    int
	Height float32
	Sex    bool
}

func main() {
	var h Human
	h = Human{
		Age:    28,
		Height: 182.0,
	}
	// what`s this?
	fmt.Printf("%p\n", &h)
	fmt.Printf("%v\n", h)
	fmt.Printf("%+v\n", h)
	fmt.Printf("%#v\n", h)
}
