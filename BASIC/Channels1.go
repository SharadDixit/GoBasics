//Channels allow us to pass data in go routines
//Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those values into another goroutine.
package main

import (
	"fmt"
)
func main(){
	messages:= make(chan string)              //channel created named messages
	go func() {messages <- "ping"}()          //go routine created where a "ping" messaged is send to the channel //sending
	msg:= <-messages                          //The <-channel syntax receives a value from the channel.    //receiving
	//Here we’ll receive the "ping" message we sent above and print it out.
	fmt.Println(msg)
}
//Note:- channel send data to go routines and not to functions directly
//for default channel which are unbuffered a receiver should be established


