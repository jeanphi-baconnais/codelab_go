package main

import (
	"fmt"
)

var infos = make (chan string)

func main() {
	// Defer
	defer fmt.Println("world")
	fmt.Println("hello")
}


