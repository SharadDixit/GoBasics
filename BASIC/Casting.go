//Well, all casting really means is taking an Object of one particular type and “turning it into” another Object type.
//This process is called casting a variable.
package main
// Atoi (string to int) and Itoa (int to string).
import (
	"fmt"
	"strconv"
)
func main(){
	randInt:= 5
	randFloat:=10.5
	randString:= "assasd"
	randString2:= "1234567891234567891"
	i:=float64(randInt)           //convert int to float by putting in a variable
	fmt.Println(i)
	fmt.Println(int(randFloat))      //convert float to int directly and then printing
	newInt, _:=strconv.ParseInt(randString,0,64)        //string to int with string name first then base as 0 and then 64bit
	fmt.Println(newInt)
	newFloat, _:=strconv.ParseInt(randString2,0,64)            //string to float with string name first then 64bit
	fmt.Println(newFloat)
}

