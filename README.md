# code-cleanup

## Motivation
Golang code-bases filled with `//TODO: remove after X amount of time` or `//TODO: remove after X happens` is a tech debt that takes away from everyone's time. It's also not the most enjoyable task to keep on looking for TODOs to clean up the code.. if it's ever done. This tool aims to clean up the code automatically and create a pull request with the changes without having to plan for code clean ups after X amount of time.

## How to use it
For blocks that you'd like to remove after a while, surround them with comments // > Start clean up and // > End clean up. Everything between those two comments would be removed.
```
func main (){
	// > Start clean up
	fmt.Println("This part should be cleaned up")
	// > End clean up
	fmt.Println("This part should not be cleaned up")
}
```
