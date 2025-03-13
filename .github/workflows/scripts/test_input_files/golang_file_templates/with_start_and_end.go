package golangfiletemplates

import "fmt"

func withStartAndEnd() {
	// > Start clean up
	fmt.Println("This part should be cleaned up")
	// > End clean up
	fmt.Println("This part should not be cleaned up")
}
