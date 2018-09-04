package main

import (
	"fmt"
)
func main(){
defer printtwo()
printone()
}
func printone(){fmt.Println(1)}
func printtwo(){fmt.Println(2)}
//using exceptions and catching exceptions
//Without defer output would have been  2 and 1 but our output is 1 and 2
//This happens because defer is gonna execute a function after the enclosing function that is even though printtwo is first but still printone gets executes first
//defer used to keep functions in a logical way and for clean up function last function
//open a file and then close it then defer can be used