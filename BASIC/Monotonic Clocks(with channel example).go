//Operating systems provide both a “wall clock,” which is subject to changes for clock synchronization, and a “monotonic clock,” which is not.
//the general rule is that the wall clock is for telling time and the monotonic clock is for measuring time.
//in this package the Time returned by time.Now contains both a wall clock reading and a monotonic clock reading;
// later time-telling operations use the wall clock reading,
// but later time-measuring operations, specifically comparisons and subtractions, use the monotonic clock reading.
package main
//Other idioms, such as time.Since(start), time.Until(deadline), and time.Now().Before(deadline), are similarly robust against wall clock resets.
//clocks will be used in main function for finding out the execution time
import (
	"fmt"
	"time"
)
func main(){
chanel := make(chan string)
		start:=time.Now()                //time now presently
			fmt.Println(start)            //prints start
        go laborArrive(chanel)
		go laborDepart(chanel)
		time.Sleep(time.Second*2)
fmt.Println(<-chanel)
	start1:=time.Now()                   //time now presently
	fmt.Println(start1)                  //prints start1
	elapsed:=start1.Sub(start)  //finds the time difference between start and start1
	elapsed1:=time.Since(start)  //finds the time since start
	fmt.Println(elapsed)   //elapsed is the execution time for the difference between start1 and start location
	fmt.Println(elapsed1)   //prints the time since start
	}
func laborArrive(chanel chan string ){
	    S:="Labor 1"
		chanel<-S
		fmt.Println(S,"is here")
		time.Sleep(time.Millisecond*10)
		}
func laborDepart(chanel chan string){
	H:=<-chanel
	fmt.Println(H,"is going")
	time.Sleep(time.Millisecond*10)
	chanel<-"Shifts Over"

}



