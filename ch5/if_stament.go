package main

import "fmt"

func main() {
	a, b := 10, true
	if a > 5 {
		fmt.Println("a > 5")
	} else if b {
		fmt.Println("a < 5")
	} else {
		fmt.Println("a == 5")
	}
	// 常见用法
	var name string
	var age int
	fmt.Println("plz input your name & age.")
	// 控制台输入格式 lx 28
	fmt.Scan(&name, &age)
	// 注意这个语句中 c 的作用域
	if c := age; c > 25 {
		fmt.Println("Too old!")
	} else if c >= 18 && c < 25 {
		fmt.Printf("you`re still young with age: %d", c)
	} else {
		fmt.Println("No way! ", c)
	}

	// 外部无法访问.
	// fmt.Println(c)
}
