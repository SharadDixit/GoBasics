//Structs will provide us a way define our own data type
package main

import "fmt"
func main(){
rect1:= rectangle{leftX:0,topY:2,height:10,width:20}
rect2:=rectangle{1,2,5,25}     //define anyway that is with variables or without variables
fmt.Println(rect1.width,rect2.width)
fmt.Println("Area of rectangle 1=",rect1.area())
	fmt.Println("Area of rectangle 2=",rect2.area())

}
type rectangle struct {
	leftX float64
	topY float64
	height float64
	width float64
}
func (rect *rectangle)area()float64 {              //for defining for a struct- first receiver(So that it knows what it is attached to) so that it works with rectangle
                                                  // then function name(also can parse value if you want) and then return data type
	finalarea:= 0.0
	finalarea= rect.width*rect.height
	return finalarea

}
