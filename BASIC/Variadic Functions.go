//Variadic functions: We can pass as many arguments as we want in the functions
package main

import (
	"fmt"
)
func main(){
	add1(3,4,5,6,7,8,9)

}
func add1(nums ...int){        //nums is a automatically generated slice
	var total int
	var i int
	var total1 int
	i=len(nums)
	fmt.Println(i)
	for j:=0;j<i ;j++  {                      //One way of computing the sum
		total1=total1+nums[j]
	}
fmt.Println(total1)                           //Another way of computing the sum
	for _, n:=range(nums){
		total=total+n

	}
	fmt.Println(total)
}

