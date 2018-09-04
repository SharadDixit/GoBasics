package main

import (
	"fmt"
	"strconv"
	"strings"
)
func main() {
	aString := "012345678"   //ASCII
	aa := len(aString)       //Length of String
	fmt.Println(aa)
	a := make([]int, aa)        //Counting the length of string for the loop
	for i := 0; i < len(aString); i++ {

		xa := int(aString[i])            //Conversion of each ASCII digit by digit into decimal
		a[i] = xa                        //Storing digit by digit in an array (integer array)
	}
	fmt.Println(a)         //printing full array
	//This step is to convert int array to string array so that we can use the join function
	var IDs []string       //For joining all the values of the string
	for i := 0; i < len(a); i++ {
		IDs = append(IDs, strconv.Itoa(a[i]))      //Appending all the values of the string by and converting values again to string from int
	}
	x := strings.Join(IDs, "")     //joining the value completely like 1 2  is 12
	fmt.Println(x)
	//Converting whole string to int is below
	newInt, _ := strconv.ParseInt(x, 0, 64)
	fmt.Println(newInt)
}

