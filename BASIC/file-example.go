package main

import (
	"os"
	"fmt"
	"io/ioutil"

	"strconv"
)
func main(){
	f,err:=os.Open("hello.txt")
	//NOTE: func Open(name string)(file *file,err error)
	// It is of os package
	// So the above function name is Open and the string name(file name) it obtains is "hello.txt" and
	// then it returns the address of the file(*file) and error for the file
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(f)
	defer f.Close()
bs,err:=ioutil.ReadAll(f)
//func ReadAll(r io.reader)([]byte, error)
//It is of ioutill package function name is ReadAll and it receives an io.reader file(which is basically *file)
//and then it returns bytes of read data([]byte) and error
if err != nil{
	fmt.Println(err)
	}
	fmt.Println(bs)
	s:= string(bs)
fmt.Println(s)
sa,_:=strconv.ParseInt(s,0,64)
fmt.Println(sa)

}

