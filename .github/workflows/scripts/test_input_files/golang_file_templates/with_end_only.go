package golangfiletemplates

import "fmt"

func withEndOnly() {
	fmt.Println("This should not be removed")
	// > End clean up
	fmt.Println("This should not be removed")
}
