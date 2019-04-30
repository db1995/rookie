package main

import (
	"fmt"
	"math"
)

func main() {
	// 表达式 T(v) 将值 v 转换为类型 T
	// Go 在不同类型的项之间赋值时需要显式转换
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	// 更加简单的形式：
	// i := 42
	// f := float64(i)
	// u := uint(f)
	fmt.Println(x, y, z)
}
