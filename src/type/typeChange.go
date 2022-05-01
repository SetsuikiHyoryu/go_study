package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	description := "Shamare is" + "138cm"
	fmt.Println(description)
	// 直接连接字符串与数值会报错
	// description02 := "Shamare is " + 138 + "cm"

	// 整数型与浮点型也不能混合使用
	pi := math.Pi
	times := 2
	// 报错
	// fmt.Println(pi * times)
	fmt.Println(pi, times)

	// 类形转换使用目标类型去包裏
	// 从浮点型转换为整数型时，小数点之后的部分会被截断而不是舍入
	fmt.Println(int(pi))

	// rune, bype 可以转换成 string，输出结果为该值对应的字符
	var piItme rune = 960
	var bang byte = 33
	fmt.Println(string(piItme), string(bang))
	// 如果需要直接输出字符串型的数字，需要使用 strconv 包的 Itoa 函数
	// Itoa: Integer to ASCII
	description02 := "Shamare is " + strconv.Itoa(138) + "cm"
	fmt.Println(description02)

	// strconv 包中亦有 Atoi 函数，但需要错误处理
	text, exception := strconv.Atoi("10")

	if exception != nil {
		fmt.Println(exception.Error())
		return
	}

	fmt.Println(text)
}
