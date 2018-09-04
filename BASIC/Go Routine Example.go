package main

import "fmt"
func main(){
f("direct")
go f("go routine")
go func(msg string){        //going routine to run a single time
	fmt.Println(msg)
	}("going")
fmt.Scanln()
fmt.Println("done")
	}
func f(from string){
	for i:=0;i<10 ;i++  {
fmt.Println(from,";",i)
	}
}
//routine of going and go routine
