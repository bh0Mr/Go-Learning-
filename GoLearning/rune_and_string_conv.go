package main

import "fmt"

func main() {

	book := "The Go language"

	runes := []rune(book)
	fmt.Printf("List of words in runes: %v\n", runes)
	
	backToBook := string(runes)
	fmt.Printf("ConVerted back to string: %s\n", backToBook)
}
