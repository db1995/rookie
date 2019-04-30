package _package

import (
	"fmt"
	"math"
)

func main() {
	// 在导入一个包时，只能引用其中大写字母开头的名字，小写字母开头的名字在该包外均无法访问
	// fmt.Println(math.pi) // error!
	fmt.Println(math.Pi)
}
