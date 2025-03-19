package main

import "fmt"

func main() {
	// > Start clean up at 2042-01-01
	fmt.Println("This part should not be cleaned up yet")
	// > Start clean up at 2023-01-01
	fmt.Println("This part should also be cleaned up")
	// > End clean up at 2023-01-01
	// > End clean up at 2042-01-01
	fmt.Println("This part should not be cleaned up")
}
