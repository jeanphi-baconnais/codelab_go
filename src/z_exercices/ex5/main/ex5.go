package main

import (
	"fmt"
)

func main() {
	// Defer
	defer fmt.Println("world")
	fmt.Println("hello")
}


