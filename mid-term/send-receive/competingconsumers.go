package main

import (
	"fmt"
	"time"
)

func sender(data string, c chan string) {
	// TODO: loop through all characters from the data.
	// Then, print "Sent: ->  X" message and send the data.
}

func receiver(receiverId string, c chan string) {
	// TODO: receive the message and print "{receiverId}  received: <-  X".
	// Then wait for one second before fetching for the next message.
}

func main() {
	var c chan string = make(chan string)
	fmt.Println("Press ENTER to exit!")

	// TODO
	// launch a Goroutine thread for the sender with "ABCDEFG" data.
	// launch a Goroutine thread for the receiver 1 with "receiver_1" as the id.
	// launch a Goroutine thread for the receiver 2 with "receiver_2" as the id.

	var input string
	fmt.Scanln(&input)
}
