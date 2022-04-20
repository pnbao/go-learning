package main

import (
	"fmt"
)

func increment() int {
	x := 0
	x++
	return x
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")

	var hello string
	hello = "Hello"
	fmt.Println(hello)

	world := "World"
	fmt.Println(hello, world)

	fmt.Println(increment())
	fmt.Println(increment())
}
