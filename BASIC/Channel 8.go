// Basic sends and receives on channels are blocking.
// However, we can use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects.
package main

import "fmt"

func main(){
	messages:=make(chan string)
	signal:=make(chan string)

	select {
	case msg:=<-messages:                 //No need of defining msg as operator can be defined
		//messages is the channel, so if the channel has received a msg to check we need to put it in a variable
		fmt.Println("Message Received",msg)
	default:
		fmt.Println("No Message Received")

	}
	sig:="Hello"                   //signal needs to defined otherwise in the case operator cannot be applied
	select {
	//signal is the channel, so to sent a msg just put it in the channel
	case signal<-sig:
		fmt.Println("Signal Sent",sig)
	default:
		fmt.Println("signal not sent")

	}
	//multi non blocking channel operation
	select {
	case msg:=<-messages:
		fmt.Println("message received",msg)
	case sig:=<-signal:
		fmt.Println("signal received",sig)
	default:
		fmt.Println("msg and signal not received")

	}

}

