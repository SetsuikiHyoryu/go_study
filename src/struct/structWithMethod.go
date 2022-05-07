package main

import "fmt"

func main() {
	shamare := operator{name: "Shamare", height: 138}
	fmt.Println(shamare.getHeight())
}

type operator struct {
	name   string
	height int
}

func (unit operator) getHeight() int {
	return unit.height
}
