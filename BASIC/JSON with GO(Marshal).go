//https://mholt.github.io/json-to-go/      (FOR JSON TO GO CONVERSION LINK)
//Go standard library includes a package for encoding JSON.
//With this package we can JSON objects to Go struct types and convert data between the two.
package main
//JSON package only encodes the public fields of a struct
	import (
	"fmt"
	"encoding/json"    //for marshal package import
)

type person struct{
	First string   //First :   alwasys first letter CAPS  that is "F" caps and not down case
	Last int
}
func main(){
p1:= person{First:"James",Last: 22}
p2:= person{"Marc", 32 }
fmt.Println(p1)
xp:=[]person{p1,p2}
fmt.Println("Go slice is:",xp)
fmt.Printf("Go slice with field names: %+v\n",xp)  //to add field names of struct then printf function with %+v
bs, err:=json.Marshal(xp)   //bs json generated will be string of bytes, so to view json format we need to convert it into string
	if err!=nil {
fmt.Println(err)
	}
	fmt.Println("JSON in bytes is :",bs)   //prints string of bytes
fmt.Println("JSON is :",string(bs))
	}

