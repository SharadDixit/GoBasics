//Range:to iterate over values received from a channel.
package main

import "fmt"
func main(){
	queue:=make(chan string,2)
	queue<-"Hello"
	queue<-"Hi"
	close(queue)
	fmt.Println(queue) //cannot excess closed channel, so use range in order to obtain information from closed channel
	for elem :=range queue{
		fmt.Println(elem)   //Now we are to access information out of closed channel
	}

	}
