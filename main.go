package main

import (
	"fmt"
)

var infos = make (chan string)

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}


