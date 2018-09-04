package main
//importing strings for all the various functions on the string; whereas string is a data type
import (
	"fmt"
"strings"
	)
func main(){
sampstring:= "Hello World"
fmt.Println(strings.Contains(sampstring,"lo"))                 //Does the string contains lo
fmt.Println(strings.Index(sampstring,"lo"))                    //Index when lo starts contains   if not present then output as -1
fmt.Println(strings.Count(sampstring,"l"))                     //Counts no. of l in the string
fmt.Println(strings.Replace(sampstring,"l","x",3))      //Replace l with x 3 times           n if is greater than present in the string then replaces all "l"
}


