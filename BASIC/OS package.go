//Package os provides a platform-independent interface to operating system functionality.
// The design is Unix-like, although the error handling is Go-like; failing calls return values of type error rather than error numbers.
// Often, more information is available within the error.
// For example, if a call that takes a file name fails, such as Open or Stat, the error will include the failing file name when printed and will be of type *PathError, which may be unpacked for more information.
package main

import (
	"os"
	"fmt"
	"log"
	"io/ioutil"
)
func main(){
	file, err:=os.Create("Sample.txt")           //creates a file named sample text        //os package needed
	if err!=nil {                               //output any types of error if occurs, not needed, log package needed
		log.Fatal(err)
	}
	file.WriteString("Randommmmm TEeEEEXXXTTT")                //Writes a string on the file
	file.Close()                                                  //close our file
	stream, err := ioutil.ReadFile("Sample.txt")          //open our file
	if err!=nil {                                   //error check again
		log.Fatal(err)
	}
	readString:= string(stream)                //to read the string in the file, converts stream to string
	fmt.Println(readString)                    //prints the string


}

