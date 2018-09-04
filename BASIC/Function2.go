package main

import (
	"fmt"
)
func main(){
	num1,num2:=next2value(5)
	fmt.Println(num1,num2)
}
func next2value(numbers int)(int,int){
	return numbers+1,numbers+2
}

