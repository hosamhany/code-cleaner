package golangfiletemplates

import "fmt"

func withStartOnly() {
	// > Start clean up
	fmt.Println("This part should not be cleaned up")
	fmt.Println("This part should not be cleaned up")
}
