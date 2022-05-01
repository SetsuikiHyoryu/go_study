package main

import "fmt"

func main() {
	count := 0

	visitCount := func() int {
		return count
	}

	fmt.Println(visitCount()) // 0

	count++
	fmt.Println(visitCount()) // 1
}
