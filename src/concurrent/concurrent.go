package main

import (
	"fmt"
	"strconv"
)

// returnnig 6 values in the channel
func show(str string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- str + "(" + strconv.Itoa(i) + ")"
	}
	channel <- "Executed! " + str
}

func main() {

	// register channels
	channel1 := make(chan string)
	channel2 := make(chan string)
	// lauch function goroutine
	go show("A", channel1)
	go show("B", channel2)

	// retrieving the channels
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
}
