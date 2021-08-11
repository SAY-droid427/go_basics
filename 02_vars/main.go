package main

import (
	"fmt"
)

// name:="Sayani"
// This can be used only inside the functions

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int 16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// runt - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	var name string = "Brad"
	var age int = 37
	size := 1.3
	fmt.Println(name, age)
	fmt.Printf("%T\n", age)
	fmt.Println(size)

	name1, email := "Sayani", "harrypotter@gmail.com"
	fmt.Println(name1, email)
}
