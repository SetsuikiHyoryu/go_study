package main

import "fmt"

func main() {
	operators := []string{"Shamare", "Suzuran"}
	dump("operators", operators)
	dump("operators[:1]", operators[:1])
}

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v\n", label, len(slice), cap(slice))
	fmt.Println(slice)
}
