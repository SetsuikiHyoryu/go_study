package main

import "fmt"

func main() {
	// 任何整数类型都可以使用 `%c` 打印字符，但是 `rune` 可以指明该数值是用来表示一个字符
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	// 打印出的是 code point
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	// 打印出的是字符
	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang)
}
