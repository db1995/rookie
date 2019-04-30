package variable

import "fmt"

// var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后
// var 语句可以出现在包或函数级别
var i, j int = 1, 2

func main() {
	// 变量声明可以包含初始值，每个变量对应一个
	// 如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型
	var c, python, java = true, false, "no!"
	// 在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明
	// 函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用
	k := 3

	fmt.Println(i, j, k, c, python, java)
}
