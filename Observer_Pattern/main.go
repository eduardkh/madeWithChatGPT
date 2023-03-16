// main.go

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

// Observer is the interface that defines the update method
type Observer interface {
	Update()
}

// Subject is the object that notifies its observers of changes
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

// TextFile is a concrete subject that watches for changes in a text file
type TextFile struct {
	observers []Observer
	filePath  string
	lastData  string
}

func (t *TextFile) RegisterObserver(observer Observer) {
	t.observers = append(t.observers, observer)
}

func (t *TextFile) RemoveObserver(observer Observer) {
	for i, o := range t.observers {
		if o == observer {
			t.observers = append(t.observers[:i], t.observers[i+1:]...)
			break
		}
	}
}

func (t *TextFile) NotifyObservers() {
	for _, observer := range t.observers {
		observer.Update()
	}
}

func (t *TextFile) Watch() {
	for {
		// Read the contents of the file
		data, err := ioutil.ReadFile(t.filePath)
		if err != nil {
			log.Println("Error reading file:", err)
		}

		// Notify the observers if the contents of the file have changed
		if string(data) != t.lastData {
			t.lastData = string(data)
			log.Printf("File %s has changed at %s\n", t.filePath, time.Now().Format(time.RFC3339))
			t.NotifyObservers()
		}

		// Wait for some time before checking the file again
		time.Sleep(1 * time.Second)
	}
}

// ObserverImpl is a concrete observer that reacts to changes in the text file
type ObserverImpl struct {
	name string
}

func (o *ObserverImpl) Update() {
	log.Printf("%s: The text file has changed!\n", o.name)
}

func main() {
	// Create a new text file subject
	textFile := &TextFile{filePath: "data.txt"}

	// Create four new observers
	observer1 := &ObserverImpl{name: "Observer1"}
	observer2 := &ObserverImpl{name: "Observer2"}
	observer3 := &ObserverImpl{name: "Observer3"}
	observer4 := &ObserverImpl{name: "Observer4"}

	// Register the observers as observers of the text file
	textFile.RegisterObserver(observer1)
	textFile.RegisterObserver(observer2)
	textFile.RegisterObserver(observer3)
	textFile.RegisterObserver(observer4)

	// Watch the text file for changes
	go textFile.Watch()

	// Wait for input from the user to exit
	var input string
	log.Println("Press any key to exit...")
	_, _ = fmt.Scanln(&input)
}
