package main

import (
	"fmt"
	"strings"
	"sort"
)
func main(){
	csvstring := "1,2,3,4,5,6"
	fmt.Println(strings.Split(csvstring,","))      //Splits the string in a list with seperation
	listofletter := []string{"c","a","b"}
	sort.Strings(listofletter)                          //sorts the string in alphabetical order
	fmt.Println("Letters=",listofletter)
	listofnums := []string{"1","2","3"}
	strings.Join(listofnums,",")                  //Joins the numbers
	fmt.Println(listofnums)

}

