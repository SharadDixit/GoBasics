package main

import (
	"fmt"
)
func main(){
dempanic()
}
func dempanic(){
	defer func(){
		fmt.Println(recover())                //panic is a type of exception which is catched by recover and so whatever we want to send
	}()
	i:=11000                                  //hit an exception that we want to send to recover
	panic(i)
}

