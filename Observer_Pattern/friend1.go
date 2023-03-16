// friend1.go

package main

import "fmt"

// Friend1 is one of the concrete observers
type Friend1 struct{}

func (f *Friend1) update(buttonPressed bool) {
	if buttonPressed {
		fmt.Println("Friend1: Yay! Let's play!")
	}
}
