package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "one"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("received:", msg1)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	select {
	case msg2 := <-c2:
		fmt.Println("received:", msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
