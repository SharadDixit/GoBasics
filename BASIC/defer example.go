package main

import "fmt"
func main(){
fmt.Println(safedivision(3,0))
fmt.Println(safedivision(6,3))

}
func safedivision(num1, num2 int)int{
	defer func(){
		recover()                       //recover catches an error if occurs - 3/0 is an error- panic error
	fmt.Println(recover())          //if you want to see the error that is nil
		}()
	solution:=num1/num2
	return solution
}
