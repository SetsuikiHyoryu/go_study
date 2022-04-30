package main

import "fmt"

func main() {
	// 字符字面值使用 `''` 包裏
	// 不指定类型的话会被推断为 rune
	grade := 'A'
	fmt.Printf("grade: %T %v %c", grade, grade, grade)
}
