package main

import "fmt"

// 切片可以用内建函数 make 来创建，这也是你创建动态数组的方式
// make 函数会分配一个元素为零值的数组并返回一个引用了它的切片
// 要指定它的容量，需向 make 传入第三个参数
func main() {
	a := make([]int, 5)
	printASlice("a", a)

	b := make([]int, 0, 5)
	printASlice("b", b)

	c := b[:2]
	printASlice("c", c)

	d := c[2:5]
	printASlice("d", d)
}

func printASlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
