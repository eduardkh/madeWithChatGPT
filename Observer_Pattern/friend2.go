// friend2.go

package main

import "fmt"

// Friend2 is another concrete observer
type Friend2 struct{}

func (f *Friend2) update(buttonPressed bool) {
	if buttonPressed {
		fmt.Println("Friend2: I'm coming too!")
	}
}
