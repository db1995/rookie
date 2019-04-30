package main

import "fmt"

type Room struct {
	Lat, Long float64
}

// 映射的文法与结构体相似，不过必须有键名
var r = map[string]Room{
	// 若顶级类型只是一个类型名，你可以在文法的元素中省略它（可以把下面两个Room去掉）
	"Bell Labs": Room{
		40.68433, -74.39967,
	},
	"Google": Room{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(r)
}
