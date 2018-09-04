package main

import (
	"fmt"
)
func main(){
	listofnumber:=[]float64{1,2,3,4,5}                     //array declaration  not defining size as if defines then also define in function where value is parsed
	fmt.Println("Sum=",addthemup(listofnumber))         //value parsing towards the function
}
func addthemup(numbers []float64) float64{                 //funtion declaration   first- parsing value with data type  Second- return data type
	sum :=0.0
	for _, val:=range numbers{
		sum = sum + val

	}
	return sum
}
