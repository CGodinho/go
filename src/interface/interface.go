package main

import (
	"fmt"
	"reflect"
)

// Whislte type definition
type Whistle string

// MakeSound method for Whistle
func (w Whistle) Sound() {
	fmt.Println(w)
}

func (w *Whistle) Change() {
	*w = "Bzuuu!"

}

// horn type definition
type Horn struct {
	message string
	id      int32
}

// MakeSound method for Horn
func (h Horn) Sound() {
	fmt.Println(h.message)
}

func (h *Horn) Change() {
	h.message = "Bqooo!"
}

func (h Horn) showId() {
	fmt.Println(h.id)
}

type Noise interface {
	Sound()
	Change()
}

func main() {

	// var is declared for the interface
	var toy Noise

	// The interface may receive any value that satisfies its implementation
	w := Whistle("Priii")
	// using a pointer refence due to change method
	toy = &w
	fmt.Println(reflect.TypeOf(toy))
	toy.Sound()
	toy.Change()
	toy.Sound()

	h := Horn{message: "Honnn", id: 3}
	toy = &h
	fmt.Println(reflect.TypeOf(toy))
	toy.Sound()
	toy.Change()
	toy.Sound()

	backHorn := (*toy).(Horn)
	//backHorn.showId()
}
