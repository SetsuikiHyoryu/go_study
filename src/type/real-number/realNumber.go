package main

import (
	"fmt"
	"math"
)

func main() {
    // 浮点型格式化动词
    floatSample := 1.0 / 3

    fmt.Printf("v: %v\n", floatSample)
    fmt.Printf("f: %f\n", floatSample)
    fmt.Printf(".3f: %.3f\n", floatSample)

    // 宽度.精度。此处意为宽度 4，精度为保留 2 位
    // 宽度大于数字长度时，左边填充空格；不指定宽度时，按实际长度显示
    // 如果要填充零，可以写作 `%04.2f`
    fmt.Printf("4.2f: %4.2f\n", floatSample)
    
    // 比较浮点型
    piggyBank := 0.1
    piggyBank += 0.2
    fmt.Println(piggyBank == 0.3) // false
    // 正确比较方式
    fmt.Println(math.Abs(piggyBank-0.3) < 0.0001)
}