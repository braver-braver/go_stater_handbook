package main

import "fmt"

// Human 结构体
type Human struct {
	Age    int
	Height float32
	Sex    bool
}

func (h *Human) Birthday() string {
	return fmt.Sprintf("%d", h.Age)
}

func (h *Human) getHeight() string {
	return fmt.Sprintf("%f", h.Height)
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
	fmt.Printf("%s\n", h.Birthday())
	fmt.Printf("%s", h.getHeight())
}
