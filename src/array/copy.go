package main

import "fmt"

func main() {
	operators := [...]string{"Shamare", "Suzuran", "April"}
	operators02 := operators

	operators[2] = "Rope"

	fmt.Println(operators)
	fmt.Println(operators02)
}
