package _package

//每个 Go 程序都是由包构成的
//程序从 main 包开始运行

import (
	"fmt"
	"math/rand"
)

func main() {
	// 按照约定，包名与导入路径的最后一个元素一致。如 "math/rand" 包中的源码均以 package rand 语句开始
	// 此程序的运行环境是固定的，因此 rand.Intn 总是会返回相同的数字
	// 要得到不同的数字，需为生成器提供不同的种子数，参见 rand.Seed
	fmt.Println("My favorite number is", rand.Intn(10))
}
