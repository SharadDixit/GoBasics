//Channel Synchroniztion: to synchronize execution across goroutines.
//Example of using a blocking receive to wait for a goroutine to finish.

package main

import (
	"fmt"
	"time"
)
func main(){
done:= make(chan bool)                    //create a channel named "done"
go worker(done)                           //Start a worker goroutine, giving it the channel to notify on //now the function starts
done<-true //If you removed the <- done line from this program, the program would exit before the worker even started //Block until we receive a notification from the worker on the channel.
}
func worker(done chan bool){          //receives the value from the done channel routine
//This is the function we’ll run in a goroutine. The done channel will be used to notify another goroutine that this function’s work is done.
fmt.Println("working")
time.Sleep(time.Millisecond*1000)
fmt.Println("done")
	<-done                         //Send a value to notify that we’re done.
	}
