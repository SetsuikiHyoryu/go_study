package main

import "fmt"

func main() {
	operators := [...]string{"Shamare", "Suzuran", "April"}

	// 也可以使用传统的三段式
	for index, operator := range operators {
		fmt.Println(index, operator)
	}
}
