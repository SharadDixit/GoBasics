//MISTAKE CHECK NEEDED OUtPUT NOT CORRECT

//Closing Channel:Closing a channel indicates that no more values will be sent on it.
// This can be useful to communicate completion to the channel’s receivers.
//making channels synchronous by done channel that is by notifying method, job sent and job received notification
package main

import "fmt"

func main(){
	jobs:=make(chan int,5)
	done:=make(chan bool)

	go func() {
		for  {
			j,more:=<-jobs
			if more {
				fmt.Println("received jobs",j)
				}else {
					fmt.Println("Received all jobs")
					done<-true
                    return
			}

		}
	}()
	for j:=1;j<=3 ;j++  {
		jobs<- j
		fmt.Println("sent job",j)
		}
		close(jobs)
		fmt.Println("sent all jobs")
		<-done            //notify that the job is sent

	}
