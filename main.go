package main

import "fmt"

func main() {
	// > Start clean up at 2042-01-01
	fmt.Println("This part should not be cleaned up yet")
	// > End clean up at 2042-01-01
	fmt.Println("This part should not be cleaned up")
}
