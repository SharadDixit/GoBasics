//Receiving Example
package main

import "fmt"

func main()  {
c1:= make(chan string)
go c1func(c1)
fmt.Println(<-c1)

	}
func c1func(c1 chan<-  string){
	c1<-"one"

}