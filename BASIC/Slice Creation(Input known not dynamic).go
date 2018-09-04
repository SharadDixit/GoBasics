//Slice is like an array but when we define it we leave out the size
package main

import "fmt"
//The type specification for a slice is []T, where T is the type of the elements of the slice.
//Unlike an array type, a slice type has no specified length.
//A slice literal is declared just like an array literal, except you leave out the element count


func main(){
	numslice:=[]int{1,2,3,4,5}
	numslice2 := numslice[0:4]   //4 is not the index but the total no. of elements
	fmt.Println(numslice2)
	// 5 is considered as 4 as the total no. of elements are 5(index from 0 to 4)
	fmt.Println("numslice2[1]=",numslice2[1])
	fmt.Println("numslice3[:2]=",numslice[:2])   //Starting from 0 to 1  (till 2 u want and not include 2)
	fmt.Println("numslice4[2:]=",numslice[2:])   //After  2 all you want  (include 2)

	numslice3:=make([]int,5,10) //make a slice whose elements are not known  where 5 is the size and 10 is the till what it can extend
	fmt.Println(numslice3)    // prints the full array, no need of for loop
	copy(numslice3,numslice)   //copy of elements of numslice to numslice 3
	fmt.Println(numslice3[2])
	numslice3=append(numslice3,0,-1)   //adding two elements 0 and -1 in numslice 3 position after index4 and size 5
	fmt.Println(numslice3)
	fmt.Println(numslice3[6])
}

