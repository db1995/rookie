package variable

import "fmt"

func main() {
	// 没有明确初始值的变量声明会被赋予它们的 零值
	// 数值类型为 0，
	// 布尔类型为 false，
	// 字符串为 ""（空字符串）。
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}