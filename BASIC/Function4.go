package main

import (
	"fmt"
)
func main(){
	fmt.Println(factorial(3))
}
//factorial(3)=3*factorial(2)
//factorial(2)=2*factorial(1)
//factorial(1)=1*factorial(0)
//factorial(0)=1
func factorial(num int)int{
	if num == 0 {
		return 1
	}
	num=num*factorial(num-1)
	return num
}

