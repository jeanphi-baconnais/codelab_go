package main

import (
	"fmt"
)

func main() {

	fmt.Println("Before call counting.")
	countingWithDefer()
	fmt.Println("After call counting.")

}

func countingWithDefer() {

	fmt.Println("Before counting ...")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("After couting ... ")
}
