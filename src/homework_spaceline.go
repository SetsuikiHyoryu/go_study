package main

import (
	"fmt"
	"math/rand"
)

func main() {
    fmt.Printf("%-20v", "SpaceLine")
    fmt.Printf("%5v", "Days")
    fmt.Printf("%-15v", "Trip type")
    fmt.Println("Price")
    
    fmt.Println("==================================================")
    
	spaceline := []string{"Space Adventures", "SpaceX", "Virgin Galactic"}
    
    for i := 0; i < 10; i++ {
        fmt.Printf("%-20v\n", spaceline[rand.Intn(3)])
    }
}