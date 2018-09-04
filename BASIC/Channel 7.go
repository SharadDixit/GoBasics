//Timeout: Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time.
package main
import (
"fmt"
"time"
)
func main (){
	c1:=make(chan string)  //c1 channel created
	go func() {
		time.Sleep(time.Second*2)  //executes after 2 seconds
		c1<-"Result1"              //receives value result1
	}()
	select {        //select function is to print basically whichever is ready first
	case res:=<-c1:      //res gets ready in 2 seconds and timeout1 in 1 seconds, therefore timeout 1 is printed
		fmt.Println(res)    //select case runs only the one whichever comes first and is ready first
	case <-time.After(time.Second*1):
		fmt.Println("timeout1")
		}
		c2:=make(chan string)       //For another example channel 2 is created
		go func() {
			time.Sleep(time.Second*2)
			c2<-"result2"        //receives value result 2 in 2 seconds
		}()
	select {
	case res1:=<-c2:               //res1 gets ready in two seconds whereas timeout 2 in 3 seconds, therefore result 2 gets ready first and prints
		fmt.Println(res1)
	case <-time.After(time.Second*3):
		fmt.Println("timeout2")


		}

}


