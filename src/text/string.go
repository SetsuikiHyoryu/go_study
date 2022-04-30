package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	operator := "巫恋"

	// len 输出字节数而不是字符数
	fmt.Println(len(operator), "bytes")
	// utf8 的包带有输出字符数的方法
	fmt.Println(utf8.RuneCountInString(operator), "runes")

	// DecodeRundeInString 返回字符串中的第一个字符和它的字节数
	runeItem, size := utf8.DecodeRuneInString(operator)
	fmt.Printf("first rune: %c, %v bytes\n", runeItem, size)

	println("===")

	// range 返回的 index 是以根据字节数计算的
	for index, item := range operator {
		fmt.Println(index, item)
		fmt.Printf("%v, %c\n", index, item)
		fmt.Println("===")
	}
}
