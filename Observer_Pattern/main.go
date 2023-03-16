// main.go

package main

import (
	"fmt"
)

func main() {
	// Create a new text file subject
	textFile := &TextFile{filePath: "data.txt"}

	// Create a new toy subject
	toy := &Toy{}

	// Create two friends who will observe the toy
	friend1 := &Friend1{}
	friend2 := &Friend2{}

	// Register the friends as observers of the toy
	toy.registerObserver(friend1)
	toy.registerObserver(friend2)

	// Register the toy as an observer of the text file
	textFile.registerObserver(toy)

	// Watch the text file for changes
	go textFile.watch()

	// Wait for input from the user to exit
	var input string
	fmt.Scanln(&input)
}
