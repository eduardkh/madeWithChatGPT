// toy.go

package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

// Observer is the interface that defines the update method
type Observer interface {
	update(buttonPressed bool)
}

// Subject is the object that notifies its observers of changes
type Subject interface {
	registerObserver(observer Observer)
	removeObserver(observer Observer)
	notifyObservers()
}

// Toy is the concrete subject in this example
type Toy struct {
	observers     []Observer
	buttonPressed bool
}

func (t *Toy) registerObserver(observer Observer) {
	t.observers = append(t.observers, observer)
}

func (t *Toy) removeObserver(observer Observer) {
	for i, o := range t.observers {
		if o == observer {
			t.observers = append(t.observers[:i], t.observers[i+1:]...)
			break
		}
	}
}

func (t *Toy) notifyObservers() {
	for _, observer := range t.observers {
		observer.update(t.buttonPressed)
	}
}

func (t *Toy) pressButton() {
	t.buttonPressed = true
	t.notifyObservers()
}

// Update is the method that implements the Observer interface
func (t *Toy) update(buttonPressed bool) {
	// The toy does not observe any other objects in this example,
	// so this method does not need to do anything
}

// TextFile is a concrete subject that watches for changes in a text file
type TextFile struct {
	observers []Observer
	filePath  string
}

func (t *TextFile) registerObserver(observer Observer) {
	t.observers = append(t.observers, observer)
}

func (t *TextFile) removeObserver(observer Observer) {
	for i, o := range t.observers {
		if o == observer {
			t.observers = append(t.observers[:i], t.observers[i+1:]...)
			break
		}
	}
}

func (t *TextFile) notifyObservers() {
	for _, observer := range t.observers {
		observer.update(true)
	}
}

func (t *TextFile) watch() {
	for {
		// Read the contents of the file
		content, err := ioutil.ReadFile(t.filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
		}

		// Notify the observers if the contents of the file have changed
		if string(content) != "" {
			t.notifyObservers()
		}

		// Wait for some time before checking the file again
		time.Sleep(1 * time.Second)
	}
}
