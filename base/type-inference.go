package main

import "fmt"

func main() {
	// 在声明一个变量而不指定其类型时（即使用不带类型的 := 语法或 var = 表达式语法），变量的类型由右值推导得出
	v := 42.6 // 修改这里使其成为不同的类型！

	// 当右值声明了类型时，新变量的类型与其相同
	var k int
	j := k // j 也是一个 int
	fmt.Printf("v is of type %T\nj is of type %T\n", v, j)

	// 不过当右边包含未指明类型的数值常量时，新变量的类型就可能是 int, float64 或 complex128 了，这取决于常量的精度
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("i: %T\nf: %T\ng: %T", i, f, g)
}
