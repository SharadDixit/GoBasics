//Map is just a collection of key value pairs, key associated with a value
package main
// The built-in function make(T, args) serves a purpose different from new(T).
// It creates slices, maps, and channels only, and it returns an initialized (not zeroed) value of type T (not *T).
// The reason for the distinction is that these three types represent, under the covers, references to data structures that must be initialized before use.
import (
	"fmt"
)
func main(){
	presidentage:= make(map[string] int)              // string mapped to integer
	presidentage["JLN"]=42                            // string JLN have value of 42
	fmt.Println(presidentage)                         // prints the complete map
	fmt.Println(presidentage["JLN"])                  // prints the value of map
    fmt.Println((len(presidentage)))                  // no. of maps
    presidentage["ManmohanSingh"]=50
    fmt.Println(len(presidentage))
    delete(presidentage,"JLN")                    //delete whichever key you want "map and the key name"
    fmt.Println(len(presidentage))                     //after deletion of key

}

//maps inside maps in the description of the video - 20:15

