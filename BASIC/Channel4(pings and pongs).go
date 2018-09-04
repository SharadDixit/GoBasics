//Channel Directions: When using channels as function parameters, you can specify if a channel is meant to only send or receive values.
// This specificity increases the type-safety of the program.
package main

import "fmt"
func main(){
pings:= make(chan string,1)
pongs:= make(chan string,1)
ping(pings,"passed msg")
pong(pings,pongs)
fmt.Println(<-pongs)
}
func ping(pings chan <-string,msg string){       //NOTE:- NO NEED OF DEFINING SENDING OR RECEIVING IN THE FUNCTION, DIRECTLY WRITE (PINGS CHAN STRING)
	pings <-msg

	}
	//func ping(pings  <- chan string,msg string){         //this will give compile time error because the function accepts the channel only for sending values
	// pings <-msg                                       //<-pings would be an error that is to try receive on the channel
	// }
func pong(pings<-chan string,pongs chan<- string){    //important line see chan position in both
	msg:=<-pings
	pongs<-msg                           //pong function accepts one channel for receiving(pings) and another channel for sending(pongs)
	}
//Each function uses pings and pongs channel differently
// receiving<-sending      if chan is on the right side then used for sending and if chan is left side then for receiving
