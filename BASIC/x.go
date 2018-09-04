//Slice is like an array but when we define it we leave out the size
package main

import "fmt"
//The type specification for a slice is []T, where T is the type of the elements of the slice.
//Unlike an array type, a slice type has no specified length.
//A slice literal is declared just like an array literal, except you leave out the element count


func main() {
	numslice := []int{1, 2, 3, 4, 5}
	numslice2 := numslice[0:5]
	fmt.Println(numslice2)
}