package main

import (
	"fmt"
)
func main(){
	num:=3
	doublenum := func() int{
		num = num*2
		return  num
	}
	fmt.Println(doublenum)                 //prints the mismatch
	fmt.Println(doublenum())               //prints the argument return

}

