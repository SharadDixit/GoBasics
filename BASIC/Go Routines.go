// A goroutine is a lightweight thread of execution.
package main

import (
	"fmt"
	"time"
)
func count(id int){
	for i:=0;i<3 ;i++  {
		fmt.Println(id,":",i)
		time.Sleep(time.Millisecond*10)
	}
}
func add(id int){
	sum:=0
	for i:=0;i<3 ;i++  {
		sum=sum+i
	}
	fmt.Println(sum)
	time.Sleep(time.Millisecond*50)
}
func main(){
	for i:=0;i<3 ;i++  {
		go count(i)         //to run the routine that is thread
	    go add(i)
		}
	time.Sleep(time.Millisecond*1000)
}

//In the program we can see that both the functions, count and add are running together simultaneously and hence go routine is achieved

