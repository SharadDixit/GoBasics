//Timers:Represents a single event in the future.
// Timers are for when you want to do something once in the future
// You tell the timer how long you want to wait, and it provides a channel that will be notified at that time.
package main

import "fmt"
import "time"
func main(){
	timer1:= time.NewTimer(time.Second*2)
	<-timer1.C              //a channel is sent values
	fmt.Println("Timer 1 expired")
	//can also use time.sleep but one property of timer is that a timer can also be stopped
	timer2:=time.NewTimer(time.Second*5)
	go func() {
		<-timer2.C              //channel is sent values
		fmt.Println("Timer 2 Expired")
	}()
	stop2:=timer2.Stop()
	if stop2==true{                  //can also be written as if stop{}
		fmt.Println("Timer2 Stop")
	}
}