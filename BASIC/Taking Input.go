package main

import (
	"fmt"
	"bufio"

	"os"

)
//The Scan function is part of "fmt", so make sure that package is imported.
// Fmt also comes with Scanf (for specifying string formatting) and Scanln (for scanning until the end of the line).
// The Scan functions store to a pointer variable.
//The scan function itself returns the number of successfully scanned items and if necessary, an error (in that order).
// It is good practice to error check when using the Scan function.
//The Scan functions are used for splitting space-delimited tokens
func main(){
	var storagevariable string
	fmt.Println("Enter String:")
	fmt.Scan(&storagevariable)
//the reader is used to read full lines.
reader := bufio.NewReader(os.Stdin)   //newreader returns *reader file           //os package needs to be imported
	// var name string     can also be used but then remove : from below line
name, _ := reader.ReadString('\n')  //takes *reader file and (returns string and error)
fmt.Println("your name is:",name)
}
