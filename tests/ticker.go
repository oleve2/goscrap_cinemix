package main

import (
	"fmt"
	"time"
)

func main() {
	d := time.NewTicker(5 * time.Second) // Calling NewTicker method
	mychannel := make(chan bool)         // Creating channel using make keyword

	/*
		go func() { // Calling Sleep() methpod in go function
			time.Sleep(7 * time.Second)
			mychannel <- true // Setting the value of channel
		}()
	*/

	// Using for loop
	for {
		select { // Select statement
		case <-mychannel: // Case statement
			fmt.Println("Completed!")
			return
		// Case to print current time
		case tm := <-d.C:
			fmt.Println("The Current time is: ", tm)
		}
	}
}
