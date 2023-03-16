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

// Subject is the object that maintains a list of its dependents (observers) and notifies them of any changes in its state
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

// Concrete Subject: TextFile is the concrete implementation of the subject interface that maintains its own state and notifies observers when its state changes
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
			log.Printf("Subject: TextFile - File %s has changed at %s\n", t.filePath, time.Now().Format(time.RFC3339))
			t.NotifyObservers()
		}

		// Wait for some time before checking the file again
		time.Sleep(1 * time.Second)
	}
}

// Concrete Observer: ObserverImpl is the concrete implementation of the observer interface that receives notifications from the subject when its state changes
type ObserverImpl struct {
	name string
}

func (o *ObserverImpl) Update() {
	log.Printf("Observer: %s - Received notification of state change\n", o.name)
}

func main() {
	// Create a new Concrete Subject: TextFile
	textFile := &TextFile{filePath: "state.txt"}

	// Create four new Concrete Observer: ObserverImpl objects
	observer1 := &ObserverImpl{name: "Observer1"}
	observer2 := &ObserverImpl{name: "Observer2"}
	observer3 := &ObserverImpl{name: "Observer3"}
	observer4 := &ObserverImpl{name: "Observer4"}

	// Register the observers as observers of the Concrete Subject: TextFile
	textFile.RegisterObserver(observer1)
	textFile.RegisterObserver(observer2)
	textFile.RegisterObserver(observer3)
	textFile.RegisterObserver(observer4)

	// Watch the Concrete Subject: TextFile for changes
	go textFile.Watch()

	// Wait for input from the user to exit
	var input string
	log.Println("Press any key to exit...")
	_, _ = fmt.Scanln(&input)
}
