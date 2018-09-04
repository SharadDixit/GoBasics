//By default channels are unbuffered,
// Meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.
// Buffered channels accept a limited number of values without a corresponding receiver for those values.
package main

import "fmt"
func main(){
messages:=make(chan string,2)
messages<-"buffered"
messages<-"channels"
fmt.Println(<-messages)
fmt.Println(<-messages)
}
//Note: No need of a receiver as channel is buffered
