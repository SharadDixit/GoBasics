//Both marshal and unmarshal accepts bytes and give output as bytes
package main

import (
	"encoding/json"
	"fmt"
)

type persons struct {
	First string
	Last string
}
//Unmarshal parses tje JSON encoded data and then stores the result in the value pointed to by v
func main(){
j:=`[{"First":"James","Last":"bond"},{"First":"Marc","Last":"Olano"}]` //this is a json slice

xp:=[]persons{}   //we need the address in order to unmarshal
err:=json.Unmarshal([]byte(j),&xp) //wants the data as a slice of bytes
if err!=nil{
	fmt.Println(err)
}
fmt.Println("Go data:",xp)
fmt.Printf("go data is %+v",xp)

}
