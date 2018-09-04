package main
import (
	"encoding/json"
	"fmt"
)

type personss struct {
	First string `json:"f"`    //This is mapping of JSON with the data struct as without this map it will loose info about First
	Last string
}
func main(){
	j:=`[{"f":"James","Last":"bond"},{"f":"Marc","Last":"Olano"}]` //this is a json slice

	xp:=[]personss{}   //we need the address in order to unmarshal
	err:=json.Unmarshal([]byte(j),&xp) //wants the data as a slice of bytes
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("Go data:",xp)
	fmt.Printf("go data is %+v",xp)

}

