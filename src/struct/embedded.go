package main

import "fmt"

func main() {
	shamare := operator{name: "Shamare", height: 138}
	fmt.Println(shamare.getHeight())
	fmt.Println(shamare.height.getHeight())
}

type operator struct {
	name
	height
}

type name string
type height int

func (_height height) getHeight() height {
	return _height
}
