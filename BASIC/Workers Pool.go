//Workers pool: Many workers will receive the work on the jobs and send the corresponding results on results
// The primary mechanism for managing state in Go is communication over channels. We saw this for example with worker pools.
package main
//Time delay is very important for go routines to work as with that they get in a organized way.
import "fmt"
import "time"
func main(){
	jobs:=make(chan int,1000)
	results:=make(chan int,1000)

	for w:=1;w<=3 ;w++  {
		go workers(w,jobs,results)
		time.Sleep(time.Second*3)
		}
	for j:=1;j<=5 ;j++  {
		jobs<-j
		time.Sleep(time.Microsecond*1000)
		}
		close(jobs)
	for r:=1;r<=5;r++  {
		<-results
		time.Sleep(time.Microsecond*1000)
	}
}
func workers(id int,jobs<-chan int,results chan<-int  ){
	for j:= range jobs{
		fmt.Println("Worker:",id,"Started Job:",j)
		time.Sleep(time.Second*2)
		fmt.Println("Workers:",id,"Finished Job",j)
		results<-j*2
		time.Sleep(time.Second*6)

	}
}

