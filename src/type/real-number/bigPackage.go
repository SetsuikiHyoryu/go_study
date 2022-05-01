package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 指数的值如果不主动声明类型则类型为 float64
	var distance = 24e18
	fmt.Println(distance)

	// 使用 big 包
	distance02 := new(big.Int)
	distance02.SetString("24000000000000000000", 10)
	fmt.Println(distance02)

	// 常量可以无类型（untype）
	const distance03 = 24000000000000000000000000
	// 程序中的每个字面值都是常量
	// 针对字面值的常量的计算在编译阶段即完成
	fmt.Println(24000000000000000000000000 / 299792 / 86400)

}
