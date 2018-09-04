package main

import (
	"fmt"
)
func main(){
	var sum[] int   //This slice has nil size
	aa:= len(sum)
	fmt.Println(aa)  //nil output right now

	var summ[10] int
	//If slice size exceeds it automatically gives 0 to other values (see ouput)
	//slice with fixed size as if you will give size from other variable it will throw an error until dynamically created
	for i:=0;i<5;i++{
		summ[i]=i
	}
	fmt.Println(summ)

	//The commented program gives a error as slice does not take size dynamically if declared like below.
	//var summms[] int
	//a:="asdasd"
	//as:=len(a)
	//for i:=0;i<len(summms(as));i++ {
	//}

	//The program below will declare size dynamically by using make
	a:="sadhaskd"
	as:=len(a)
	summm:= make([]int,as) //where as is the size of the slice developed
	fmt.Println(len(summm))    //size given dynamically



}

