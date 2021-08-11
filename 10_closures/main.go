package main

import "fmt"

// Go language provides a special feature known as an anonymous function.
// An anonymous function can form a closure. A closure is a special type of anonymous function that references variables declared outside of the function itself.
// It is similar to accessing global variables which are available before the declaration of the function.

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}
