package main

import "fmt"

//this program does not change the global variable value of x defined in x, therefore we need to use pointers
//Because with the help of pointers we can change the value at the address location of the x variable
//func main(){
//x:=0
//changeval(x)
//fmt.Println("x=",x)
//}
//func changeval(x int){
//	x=2
//}

func main(){
x:=0
changevalnow(&x)                                 //passing reference of x (that is address)
fmt.Println("x=",x)
fmt.Println(&x)                                   //
}
func changevalnow(x *int){                       //reference to the value and not the value itself, change the value at memory address, deals with address
*x=2
}