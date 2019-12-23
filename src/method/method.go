package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

// (f Fahrenheit) is the receiver and receiver parameter
// f by convention is just a letter
// method is only possible to be defined within the same pacakge where the type
// is defined (in this case Fahrenheit)
func (f Fahrenheit) toCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (f *Fahrenheit) increment() {
	*f++
}

func main() {
	var temp1 Fahrenheit = Fahrenheit(50)
	temp2 := temp1.toCelsius()
	fmt.Println("Output is: ", temp2)
	// call needs a pointer, but it is automatically taken
	// only possible to get a reference to a variable
	temp1.increment()
	fmt.Println("Output is: ", temp1)
}
