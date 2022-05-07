package main

import "fmt"

func main() {
	operators := [...]string{"Shamare", "Suzuran"}
	operatorsSlice := operators[:]

	operatorsSlice2 := []string{"Shamare", "Suzuran"}

	fmt.Println(operatorsSlice)
	fmt.Println(operatorsSlice2)
	fmt.Printf("%T %T\n", operators, operatorsSlice2)
}
