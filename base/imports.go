package main

// 用圆括号组合了导入，比写多个import语句好
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
