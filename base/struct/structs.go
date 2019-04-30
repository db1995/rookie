package main

import "fmt"

// 一个结构体（struct）就是一组字段（field）
type Vertex struct {
	X int
	Y int
}

type OneStruct struct {
	X, Y int
}

// 结构体文法通过直接列出字段的值来新分配一个结构体
var (
	v1 = OneStruct{1, 2} // 创建一个 OneStruct 类型的结构体
	// 使用 Name: 语法可以仅列出部分字段（字段名的顺序无关）
	v2 = OneStruct{X: 1} // Y:0 被隐式地赋予
	v3 = OneStruct{}     // X:0 Y:0
	// 特殊的前缀 & 返回一个指向结构体的指针
	p1 = &OneStruct{1, 2} // 创建一个 *OneStruct 类型的结构体（指针）
)

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	/// 结构体字段使用点号来访问
	v.X = 4
	fmt.Println(v.X)

	// 结构体字段可以通过结构体指针来访问
	// 如果我们有一个指向结构体的指针 p，那么可以通过 (*p).X 来访问其字段 X
	// 也允许我们使用隐式间接引用，直接写 p.X 就可以
	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, p1, v2, v3)
}
