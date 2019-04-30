package main

import "fmt"

// Go 只有一种循环结构：for 循环
/*
	基本的 for 循环由三部分组成，它们用分号隔开：

	初始化语句：在第一次迭代前执行
	条件表达式：在每次迭代前求值
	后置语句：在每次迭代的结尾执行
	初始化语句通常为一句短变量声明，该变量声明仅在 for 语句的作用域中可见

	一旦条件表达式的布尔值为 false，循环迭代就会终止
*/
func main() {
	sumA := 0
	for i := 0; i < 10; i++ {
		sumA += i
	}
	fmt.Println(sumA)

	// 初始化语句和后置语句是可选的
	sumB := 1
	for sumB < 1000 {
		sumB += sumB
	}
	fmt.Println(sumB)

	// 此时你可以去掉分号，因为 C 的 while 在 Go 中叫做 for
	sumC := 1
	for sumC < 1000 {
		sumC += sumC
	}
	fmt.Println(sumC)

	// 如果省略循环条件，该循环就不会结束
	for {
		fmt.Println("----")
	}
}
