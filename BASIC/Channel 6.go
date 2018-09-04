//Select: Go’s select lets you wait on multiple channel operations.
// Combining goroutines and channels with select is a powerful feature of Go.
//select proceeds with the first receive that’s ready
package main

import (
	"fmt"
	"time"
)
func main() {
	c1 := make(chan string)
	c2 := make(chan string)
go func(){
	c1<-"one"
	time.Sleep(time.Millisecond*1000)         //Each channel will receive a value after some amount of time, to simulate
	}()
go func(){
	c2<-"two"
	time.Sleep(time.Millisecond*5000)          //Each channel will receive a value after some amount of time, to simulate
	//e.g. blocking RPC operations executing in concurrent goroutines. RPC: Remote procedure call
	// An RPC is initiated by the client, which sends a request message to a known remote server to execute a specified procedure with supplied parameters. The remote server sends a response to the client, and the application continues its process.
	// While the server is processing the call, the client is blocked (it waits until the server has finished processing before resuming execution),
	// Unless the client sends an asynchronous request to the server, such as an XHTTP call.
}()
	for i:=0;i<2;i++  {
		select {       //We’ll use select to await both of these values simultaneously, printing each one as it arrives.
		case msg1:=<-c1:  //selects only runs one case but since for loop is there, therefore execution is twice
		fmt.Println(msg1)
		case msg2:=<-c2:
			fmt.Println(msg2)


		}   //c1 has 1 sec wait and c2 has 5 sec wait, therefore c1 should be readied before c2 and hence one pro

	}
	}

