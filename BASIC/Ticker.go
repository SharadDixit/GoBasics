//Ticker: tickers are for when you want to do something repeatedly at regular intervals.
package main

import "fmt"
import "time"
func main(){
	//Tickers use a similar mechanism to timers: a channel that is sent values
	ticker:=time.NewTicker(time.Microsecond*500)
	go func() {
		for t:=range ticker.C{
			fmt.Println("Tick at",t)
			}

	}()
	time.Sleep(time.Millisecond*1600)
	ticker.Stop()
	fmt.Println("Ticker Stopped")
}
