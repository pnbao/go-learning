package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"  //send value to channel
		messages <- "ping2" //send value to channel
		messages <- "ping3" //send value to channel
		messages <- "ping4" //send value to channel
		messages <- "ping5" //send value to channel
		messages <- "ping6" //send value to channel
	}()

	msg := <-messages  //get value from channel
	msg2 := <-messages //get value from channel
	msg3 := <-messages //get value from channel
	fmt.Println(msg, msg2, msg3)
}
