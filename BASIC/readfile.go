package main

import (
	"fmt"
	"io/ioutil"
)
func main(){
	bs,err:=ioutil.ReadFile("Hello.txt")
	//func ReadFile(file name string)([]byte,error)
	// It is of ioutill package.
	// Receives a file by the name of string("Hello.txt")
	// Returns a bytes of read file in a slice and error

	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(bs)
	s:=string(bs)
	a:=len(bs)
fmt.Println(a)
	//bytes of slice file is converted into string
	fmt.Println(s)
}
