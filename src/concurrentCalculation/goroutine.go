package main

import (
	"fmt"
	"time"
)

func main() {
	go shamareAttack()
	suzuranAttack()
}

func shamareAttack() {
	for i := 0; i < 2; i++ {
		fmt.Println(i+1, "Shamare attacked.")
		time.Sleep(time.Second)
	}
}

func suzuranAttack() {
	for i := 0; i < 3; i++ {
		fmt.Println(i+1, "Suzuran attacked.")
		time.Sleep(time.Second)
	}
}
