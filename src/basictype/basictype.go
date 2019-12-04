package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	// unsigned int 8 bits
	var v1 uint8 = 255
	fmt.Println(v1)

	// signed int 8 bits
	var v2 int8 = 127
	fmt.Println(v2)

	// from 8 up to 64 bits
	var v3 uint64 = 255
	var v4 int64 = 127
	fmt.Printf("%d - %d\n", v3, v4)

	// floats
	var f1 float32 = 1.12121
	var f2 float64 = 5.56656
	fmt.Printf("%4.2f - %0.2f\n", f1, f2)

	// complex
	c1 := 0.867 + 0.5i // complex128
	fmt.Printf("%g\n", c1)

	// boolean
	b := true
	fmt.Printf("%t\n", b)

	// constants
	const Pi = 3.14

	// Zero values
	// Variables declared without an explicit initial value are given their zero value.
	//
	// The zero value is:
	//
	// - 0 for numeric types,
	// - false for the boolean type, and
	// - "" (the empty string) for strings.

	str1 := "hello"
	str2 := "12.23"
	fmt.Printf("%s %s\n", str1, str2)

	//convert
	i, _ := strconv.Atoi("-42")
	fmt.Printf("%d\n", i)

	s := strconv.Itoa(-42)
	fmt.Printf("%s\n", s)

	// string to types
	b, _ = strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i2, _ := strconv.ParseInt("-42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)
	fmt.Printf("%T %0.3f %d %d\n", b, f, i2, u)

	// type to string
	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-42, 16)
	s4 := strconv.FormatUint(42, 16)
	fmt.Printf("%s %s %s %s\n", s1, s2, s3, s4)

	// char -> rune
	aRune := 'G'
	fmt.Printf("Rune 1: %c; Unicode: %U; Type: %s\n", aRune, aRune, reflect.TypeOf(aRune))

}
