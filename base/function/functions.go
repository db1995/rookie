package function

import "fmt"

// 函数可以没有参数或接受多个参数
func getTen() int {
	return 10
}

// 注意类型在变量名之后
func add(x int, y int) int {
	return x + y
}

// 当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略
func sub(x, y int) int {
	return x - y
}

func main() {
	fmt.Println(getTen())
	fmt.Println(add(42, 13))
	fmt.Println(sub(23, 9))
}
