package main

import "fmt"

// 声明包级别变量
var (
	// 以大写开头的 命名表明可以在 main 包外访问.
	globalVal = "global"
	GlobalVal = "Global"
)

// 不允许,短声明语法仅限于 函数内部
// _globalVal := "global"

// 常量的声明方式
const ()

const (
	_ = iota
	// what is the values of below vars ?
	// 试着答应一下
	Monday
	Tuesday
)

func main() {
	var (
		a     int
		b     int
		c     int
		age   byte
		sex   bool
		price float64
	)
	// 变量赋值
	a = 10
	b, c = 20, a+b
	fmt.Printf(" linecode: 28 , a = %d, b = %d, c = %d", a, b, c)
	// := 声明与赋值放在同一行代码
	f := 40.0
	g := c - a

	// 使用 var 关键字声明变量, 类型为 int, 等价为
	// h := 4
	var h int = 4
	// 自动从值推断变量类型(int)
	var m = 4
	fmt.Print("linecode: 38 ", h, m)
	// 向标准输出输出 c 的值,并在结尾添加换行符
	fmt.Println("linecode: 40 ", c)
	// 观察这行的输出, go 语言变量声明后的默认值均为"零"值.
	fmt.Printf("%d %d %t %.2f\n", c, age, sex, price)
	fmt.Printf("%T %T\n", f, g)
}
